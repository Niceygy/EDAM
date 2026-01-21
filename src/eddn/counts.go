package eddn

import "time"

/*
Single uploader entry with FSDMessages
seen and timestamp when it was measured.

FSDMessages are used becuase they provide
a good approxomation for the activity of a user.

Different ones may be used in the future.
*/
type UploaderEntry struct {
	Messages  int
	Timestamp time.Time
}

const UPLOADER_COUNT_TIME time.Duration = time.Minute * 1

/*How often should UPLOADERS_PAST_HOUR be updated and saved to FTP?*/
const EDDN_CSV_BACKUP_INTERVAL time.Duration = time.Hour * 1

/*String list of all uploader IDs since last refresh*/
var UPLOADERS_SINCE_REFRESH int

/*
UploaderEntry list with the number of uploaders
and the timestamp when it was measured. Updated
every minute.
*/
var UPLOADERS_PAST_HOUR []UploaderEntry

/*
UploaderEntry list with the number of uploaders
and the timestamp when it was measured. Updated
every hour. Restored from FTP when app starts.
*/
var UPLOADERS_ALL_TIME []UploaderEntry

/*
CSV data of uploaders/hour.
Calculated by (every hour) averaging the
number of uploaders in UPLOADERS_SINCE_REFRESH
and multiplying it by 60.
*/
var CSV_FOR_FTP string

type EDStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Product string `json:"product"`
}

type EDState bool

const (
	EDStateOnline  EDState = true
	EDStateOffline EDState = false
)
