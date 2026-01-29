package eddn

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-zeromq/zmq4"
)

var UploaderChannel = make(chan EDMessageType)

/*Entrypoint. Connects to EDDN and launches all related goroutines*/
func EDDNListener() {
	restoreFromFTP(false)
	go onTheRefreshHandler()
	go csvBackupHandler()

	for {

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

			if message.Event == EDMessage_FSD || message.Event == EDMessage_Docked {
				UPLOADERS_SINCE_REFRESH++
			}

			select { //only send if there is something to recive it
			case UploaderChannel <- message.Event:
			default:
			}
		}

		log.Println("STOP")
	}
}

/*Decodes a ZLIB encoded message into a EDDNMessage struct*/
func decodeMessage(rawMessage []byte) (*EDDNMessage, error) {
	r, err := zlib.NewReader(bytes.NewReader(rawMessage))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer r.Close()

	var message EDDNMessage
	err = json.NewDecoder(r).Decode(&message)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//Find and save the event type
	var inner map[string]any
	if err := json.Unmarshal(message.Message, &inner); err != nil {
		return nil, fmt.Errorf("failed to unmarshal inner message: %w", err)
	}
	if event, ok := inner["event"]; ok {
		switch event.(string) {
		case "FSDJump":
			message.Event = EDMessage_FSD
			// log.Println("FSDJump")
		case "DockingGranted", "DockingDenied":
			message.Event = EDMessage_Docked
			// log.Println("Docked")
		default:
			message.Event = EDMessage_Other
		}
	}

	return &message, nil
}

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
