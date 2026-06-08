package eddn

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

// type EDMessageType int8

// const (
// 	EDMessage_FSD    EDMessageType = 0
// 	EDMessage_Docked EDMessageType = 1
// 	EDMessage_Other  EDMessageType = 2
// )

// type EDDNMessage struct {
// 	SchemaRef string          `json:"$schemaRef"`
// 	Header    EDDNHeader      `json:"header"`
// 	Message   json.RawMessage `json:"message"`
// 	Event     EDMessageType
// }
// type EDDNHeader struct {
// 	UploaderID       string    `json:"uploaderID"`
// 	SoftwareName     string    `json:"softwareName"`
// 	SoftwareVersion  string    `json:"softwareVersion"`
// 	GatewayTimestamp time.Time `json:"gatewayTimestamp"`
// }

var CSV_FOR_WEB string

var HIGHEST_SEEN_EDDN_COUNT int = -1
