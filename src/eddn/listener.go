package eddn

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/go-zeromq/zmq4"
)

const UPLOADER_COUNT_TIME time.Duration = time.Minute * 1
const EDDN_CSV_BACKUP_INTERVAL time.Duration = time.Hour * 1

var UPLOADERS_SINCE_REFRESH []string

type EDDNMessage struct {
	SchemaRef string          `json:"$schemaRef"`
	Header    EDDNHeader      `json:"header"`
	Message   json.RawMessage `json:"message"`
}
type EDDNHeader struct {
	UploaderID       string    `json:"uploaderID"`
	SoftwareName     string    `json:"softwareName"`
	SoftwareVersion  string    `json:"softwareVersion"`
	GatewayTimestamp time.Time `json:"gatewayTimestamp"`
}

var UploaderChannel = make(chan string)

func EDDNListener() {
	restoreFromFTP()
	go onTheRefreshHandler()
	go eddnMessageHandler()

	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	err := sub.Dial("tcp://eddn.edcd.io:9500")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = sub.SetOption(zmq4.OptionSubscribe, "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	run := true
	for run {
		//for each message
		msg, err := sub.Recv()
		if err != nil {
			fmt.Println(err.Error())
			run = false
			continue
		}

		message, err := decodeMessage(msg.Frames[0])
		if err != nil {
			fmt.Println(err.Error())
			run = false
			continue
		}

		UploaderChannel <- message.Header.UploaderID
	}
}

/*Decodes a ZLIB encoded message into a EDDNMessage struct*/
func decodeMessage(rawMessage []byte) (*EDDNMessage, error) {
	r, err := zlib.NewReader(bytes.NewReader(rawMessage))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var message EDDNMessage
	err = json.NewDecoder(r).Decode(&message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func eddnMessageHandler() {
	for {
		uploaderID := <-UploaderChannel
		if !slices.Contains(UPLOADERS_SINCE_REFRESH, uploaderID) {
			UPLOADERS_SINCE_REFRESH = append(UPLOADERS_SINCE_REFRESH, uploaderID)
		}
	}
}

func onTheRefreshHandler() {
	for {
		time.Sleep(UPLOADER_COUNT_TIME)
		EDDN_CSV_DATA = strings.Join([]string{
			EDDN_CSV_DATA,
			"\n",
			strconv.FormatInt(time.Now().Unix(), 10),
			",",
			strconv.Itoa(len(UPLOADERS_SINCE_REFRESH)),
		}, "")
		log.Println("Seen " + strconv.Itoa(len(UPLOADERS_SINCE_REFRESH)) + " Uploaders in the past 1m")
		UPLOADERS_SINCE_REFRESH = []string{}
	}
}
