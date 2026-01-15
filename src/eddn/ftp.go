package eddn

import (
	"bytes"
	"io"
	"log"
	"os"
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

func restoreFromFTP() {
	c, err := ftp.Dial(getEnvVar("FTP_ADDR")+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(getEnvVar("FTP_USERNAME"), getEnvVar("FTP_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	r, err := c.Retr("test-file.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	EDDN_CSV_DATA = string(buf)
}

func csvBackupHandler() {
	for {
		time.Sleep(EDDN_CSV_BACKUP_INTERVAL)

		c, err := ftp.Dial(getEnvVar("FTP_ADDR")+":21", ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			log.Fatal(err)
		}

		err = c.Login(getEnvVar("FTP_USERNAME"), getEnvVar("FTP_PASSWORD"))
		if err != nil {
			log.Fatal(err)
		}

		err = c.Delete(getEnvVar("FTP_FULLPATH"))

		if err != nil {
			log.Panic(err)
		}

		data := bytes.NewBufferString(EDDN_CSV_DATA)
		err = c.Stor(getEnvVar("FTP_FULLPATH"), data)
		if err != nil {
			panic(err)
		}

		c.Logout()
		c.Quit()
	}
}
