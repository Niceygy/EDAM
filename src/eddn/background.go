package eddn

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func EDDN_RecordHourlyCounts() {
	data, err := os.ReadFile(os.Getenv("HISTORICAL_DATA_FILE"))
	if err != nil {
		log.Panic(err)
	}

	CSV_FOR_WEB = string(data)

	for {
		currentCount := GetCurrentEDDNCount()
		currentUnixTime := time.Now().Unix()

		lineToAdd := "\n" + strconv.Itoa(int(currentUnixTime)) + "," + currentCount

		CSV_FOR_WEB += lineToAdd

		os.WriteFile(os.Getenv("HISTORICAL_DATA_FILE"), []byte(CSV_FOR_WEB), os.ModeAppend)

		time.Sleep(time.Hour * 1)
	}
}

func GetHighestEDDNCount() int {
	if HIGHEST_SEEN_EDDN_COUNT > 0 {
		return HIGHEST_SEEN_EDDN_COUNT
	} else {
		for line := range strings.SplitSeq(CSV_FOR_WEB, "\n") {
			if line == "\n" || len(line) < 1 {
				continue
			}
			line = strings.TrimSpace(line)
			count, err := strconv.Atoi(strings.Split(line, ",")[1])
			if err != nil {
				log.Panic(err)
			}

			if count > HIGHEST_SEEN_EDDN_COUNT {
				HIGHEST_SEEN_EDDN_COUNT = count
			}
		}
	}

	return HIGHEST_SEEN_EDDN_COUNT
}
