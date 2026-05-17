package eddn

import (
	"strconv"
	"time"

	"niceygy.net/edam/src/errors"
)

var UploaderChannel = make(chan EDMessageType)

/*Entrypoint. Connects to EDDN and launches all related goroutines*/
func EDDNListener() {
	restoreFromFile(false)
	// go onTheRefreshHandler()
	go csvBackupHandler()

	for {
		time.Sleep(UPLOADER_COUNT_TIME)

		var entry UploaderEntry
		entry.Timestamp = time.Now()
		count, err := strconv.Atoi(GetCurrentEDDNCount())
		errors.PanicIfErr(err)
		entry.Messages = count

		UPLOADERS_PAST_HOUR = append(UPLOADERS_PAST_HOUR, entry)
		UPLOADERS_SINCE_REFRESH = 0
	}
}
