package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const EDDN_CSV_FILEPATH = "static/data/messageCount.csv"

var LAST_UPDATED time.Time = time.Now()

func EDDNCsvLoop() {
	for {
		downloadEDDNCsv()
		time.Sleep(time.Hour * 1)
	}
}

func downloadEDDNCsv() {
	req, err := http.NewRequest(http.MethodGet, "https://niceygy.net/experiments/edam/data.csv", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return
	}
	defer resp.Body.Close()
	// filedata := []byte{}
	// resp.Body.Read(filedata)
	filedata, err := io.ReadAll(resp.Body)

	err = os.Remove(EDDN_CSV_FILEPATH)
	err = os.WriteFile(EDDN_CSV_FILEPATH, filedata, os.ModeAppend)
}
