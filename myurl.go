package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var urlTable = make(map[string]string)

const URLTABLEFLIE string = "data/url.csv"

func loadURLTable() {
	// create empty url.csv file if not exists
	if _, err := os.Stat(URLTABLEFLIE); errors.Is(err, os.ErrNotExist) {
		file, err2 := os.Create(URLTABLEFLIE)
		if err2 != nil {
			panic("[ERROR] failed to create url.csv, error: " + err2.Error())
		}
		defer file.Close()
	}

	file, err := os.Open(URLTABLEFLIE)
	if err != nil {
		panic("[ERROR] failed to open url.csv, error: " + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	data, err := reader.ReadAll()
	if err != nil {
		panic("[ERROR] failed to read the content of url.csv, error: " + err.Error())
	}

	for _, row := range data {
		if len(row) != 2 {
			fmt.Printf("Skipping invalid row: %v", row)
			continue
		}
		urlTable[row[0]] = row[1]
	}
}

func writeURLTable() {
	file, err := os.Create(URLTABLEFLIE)
	if err != nil {
		panic("[ERROR] failed to create url.csv, error: " + err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for url, destination := range urlTable {
		record := []string{url, destination}
		if err := writer.Write(record); err != nil {
			fmt.Print("error while writing record ({" + url + ", " + destination + "}) to url.csv: " + err.Error())
		}
	}
}

func createURL(shortened string, destination_raw string, override bool) (status string, message string) {
	if shortened == "admin" {
		return "failed", "Cannot use `admin` as entry"
	}

	destination, ok := urlValidator(destination_raw)
	if !ok {
		return "failed", destination
	}

	record, exists := urlTable[shortened]
	if !exists || (exists && override) {
		urlTable[shortened] = destination
		writeURLTable()
		return "successful", "Added: (" + shortened + " -> " + destination + ")"
	} else if exists && !override {
		return "override-confirm", "Already exists (" + shortened + " -> " + record + "), please use override"
	} else {
		return "failed", "Unknown error"
	}
}

func urlValidator(input string) (parsed string, ok bool) {
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "https://" + input
	}

	parsedURL, err := url.Parse(input)
	if err != nil {
		return "INVALID URL", false
	}

	hostPattern := `^([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}$`
	validHost := regexp.MustCompile(hostPattern).MatchString(parsedURL.Hostname())
	if !validHost {
		return "INVALID HOST: " + parsedURL.Hostname(), false
	}

	if parsedURL.Hostname() == "t.gravitycat.tw" {
		return "CANNOT USE THE HOST: t.gravitycat.tw", false
	}

	return parsedURL.String(), true
}
