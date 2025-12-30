package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const EDDN_CSV_FILEPATH = "static/data/messageCount.csv"

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

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("client: bad status code: %d\n", resp.StatusCode)
		return
	}

	filedata, err := io.ReadAll(resp.Body)

	err = os.Remove(EDDN_CSV_FILEPATH)
	err = os.WriteFile(EDDN_CSV_FILEPATH, filedata, 0644)
	if err != nil {
		fmt.Printf("client: error writing file: %s\n", err)
	}
}

func getHighestEDDNCount() int {
	data, err := os.ReadFile("static/data/messageCount.csv")

	if err != nil {
		log.Println("ERR Open static/data/messageCount.csv: " + err.Error())
	}

	stringdata := string(data)

	lines := strings.Split(stringdata, "\n")
	largest := 0

	for i := range lines {
		line := lines[i]

		if line == "" {
			break
		}

		count, err := strconv.Atoi(strings.Split(line, ",")[1])
		if err != nil {
			log.Panic(err.Error())
		}

		if count > largest {
			largest = count
		}
	}

	return largest
}

func getCurrentEDDNCount() string {
	data, err := os.ReadFile("static/data/messageCount.csv")

	if err != nil {
		log.Println("ERR Open static/data/messageCount.csv: " + err.Error())
	}

	stringdata := string(data)

	lines := strings.Split(stringdata, "\n")
	line := lines[0]

	return strings.Split(line, ",")[1]
}
