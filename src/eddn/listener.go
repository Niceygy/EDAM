package eddn

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-zeromq/zmq4"
)

type EDDNMessage struct {
	SchemaRef string          `json:"$schemaRef"`
	Header    EDDNHeader      `json:"header"`
	Message   json.RawMessage `json:"message"`
	Event     string
}
type EDDNHeader struct {
	UploaderID       string    `json:"uploaderID"`
	SoftwareName     string    `json:"softwareName"`
	SoftwareVersion  string    `json:"softwareVersion"`
	GatewayTimestamp time.Time `json:"gatewayTimestamp"`
}

// var UploaderChannel = make(chan string)

/*Entrypoint. Connects to EDDN and launches all related goroutines*/
func EDDNListener() {
	restoreFromFTP(false)
	go onTheRefreshHandler()
	// go eddnMessageHandler()
	go csvBackupHandler()

	//open the connection
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

		if message.Event == "FSDJump" {
			UPLOADERS_SINCE_REFRESH++
		}
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

	//Find and save the event type
	var inner map[string]any
	if err := json.Unmarshal(message.Message, &inner); err != nil {
		return nil, fmt.Errorf("failed to unmarshal inner message: %w", err)
	}
	if event, ok := inner["event"]; ok {
		message.Event = event.(string)
	}

	return &message, nil
}

// /*
// Manages the UPLOADERS_SINCE_REFRESH list.
// If it sees a uploaderID that isn't in the list,
// it adds it.
// */
// func eddnMessageHandler() {
// 	for {
// 		uploaderID := <-UploaderChannel
// 		if !slices.Contains(UPLOADERS_SINCE_REFRESH, uploaderID) {
// 			UPLOADERS_SINCE_REFRESH = append(UPLOADERS_SINCE_REFRESH, uploaderID)
// 		}
// 	}
// }

/*
Every UPLOADER_COUNT_TIME, updates the UPLOADERS_PAST_HOUR
with data from the UPLOADERS_SINCE_REFRESH slice and
clears it.
*/
func onTheRefreshHandler() {
	for {
		time.Sleep(UPLOADER_COUNT_TIME)

		var entry UploaderEntry
		entry.Timestamp = time.Now()
		entry.Messages = UPLOADERS_SINCE_REFRESH

		UPLOADERS_PAST_HOUR = append(UPLOADERS_PAST_HOUR, entry)
		UPLOADERS_SINCE_REFRESH = 0

		// log.Println("Seen " + strconv.Itoa(entry.Uploaders) + " in the past minute")
	}
}
