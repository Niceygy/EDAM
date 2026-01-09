package main

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-zeromq/zmq4"
)

type EDDN struct {
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

var Uploaders []string
var uploaderChannel = make(chan string)

func eRelay() {
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
		select {
		default:
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

			uploaderChannel <- message.Header.UploaderID
		}
	}
}

func decodeMessage(rawMessage []byte) (*EDDN, error) {
	r, err := zlib.NewReader(bytes.NewReader(rawMessage))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var message EDDN
	err = json.NewDecoder(r).Decode(&message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
