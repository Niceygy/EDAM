package eddn

import (
	"bytes"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

func getEnvVar(key string) string {
	value, isPresent := os.LookupEnv(key)
	if !isPresent {
		log.Panic(key + " is not present!")
		return ""
	} else {
		return value
	}
}

func openFTP() *ftp.ServerConn {
	c, err := ftp.Dial(getEnvVar("FTP_ADDRESS")+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(getEnvVar("FTP_USERNAME"), getEnvVar("FTP_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func restoreFromFTP(returnNotRestore bool) []UploaderEntry {
	conn := openFTP()
	r, err := conn.Retr(getEnvVar("FTP_FULLPATH"))
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	if (string(buf)) == "" {
		UPLOADERS_ALL_TIME = []UploaderEntry{}
		log.Println("Skipped FTP restore (no data found)")
		return []UploaderEntry{}
	}
	data := strings.Split(string(buf), "\n")

	var result []UploaderEntry

	for i := range data {
		line := data[i]

		if line == "" {
			continue
		}

		_time, err := strconv.ParseInt(strings.Split(line, ",")[0], 10, 64)
		if err != nil {
			panic(err)
		}

		var entry UploaderEntry
		entry.Timestamp = time.Unix(int64(_time), 0)
		entry.Messages, err = strconv.Atoi(strings.Split(line, ",")[1])

		if err != nil {
			panic(err)
		}

		result = append(result, entry)
	}

	log.Println("Restored from FTP")

	if returnNotRestore {
		return result
	} else {
		UPLOADERS_ALL_TIME = result
		return []UploaderEntry{}
	}

}

/*
Handles:

- Adveraging the data from UPLOADERS_PAST_HOUR,

- Putting that data into UPLOADERS_ALL_TIME,

- Making a CSV file out of UPLOADERS_ALL_TIME,

- Saving that CSV to the FTP server
*/
func csvBackupHandler() {
	for {
		time.Sleep(EDDN_CSV_BACKUP_INTERVAL)

		//average the data from UPLOADERS_PAST_HOUR

		var totalUploaders int64

		for i := range UPLOADERS_PAST_HOUR {
			entry := UPLOADERS_PAST_HOUR[i]

			totalUploaders += int64(entry.Messages)
		}

		average := float64(totalUploaders / int64(len(UPLOADERS_PAST_HOUR)))
		average = math.Round(average)
		UPLOADERS_PAST_HOUR = []UploaderEntry{}

		var entry UploaderEntry
		entry.Timestamp = time.Now()
		entry.Messages = int(average)

		// get the old CSV & update it

		var oldCSV []UploaderEntry = restoreFromFTP(true)
		var newCSV []UploaderEntry = append(oldCSV, entry)

		//convert it back to a string

		var stringCSV string

		for i := range newCSV {
			stringCSV += strings.Join([]string{
				strconv.Itoa(int(newCSV[i].Timestamp.Unix())),
				",",
				strconv.Itoa(newCSV[i].Messages),
				"\n",
			}, "")
		}
		//save it

		conn := openFTP()
		err := conn.Delete(getEnvVar("FTP_FULLPATH"))

		if err != nil {
			log.Panic(err)
		}

		CSV_FOR_FTP = stringCSV
		data := bytes.NewBufferString(stringCSV)
		err = conn.Stor(getEnvVar("FTP_FULLPATH"), data)
		if err != nil {
			panic(err)
		}

		conn.Logout()
		conn.Quit()

		log.Println("Saved to CSV. Saw " + strconv.Itoa(int(average)) + " average uploaders in the past hour")
	}
}
