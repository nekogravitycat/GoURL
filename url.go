package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var urlTable = make(map[string]string)

const URLTABLEFLIE string = "data/url.csv"

func loadURLTable() {
	file, err := os.Open(URLTABLEFLIE)
	if err != nil {
		fmt.Print("error while opening url.csv: " + err.Error())
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	data, err := reader.ReadAll()
	if err != nil {
		fmt.Print("error while reading url.csv: " + err.Error())
		return
	}

	for _, row := range data {
		urlTable[row[0]] = row[1]
	}

	fmt.Print(urlTable)
}

func writeURLTable() {
	file, err := os.Create(URLTABLEFLIE)
	if err != nil {
		fmt.Print("error while creating url.csv: " + err.Error())
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for url, destination := range urlTable {
		if err := writer.Write([]string{url, destination}); err != nil {
			fmt.Print("error while writing record ({" + url + ", " + destination + "}) to url.csv: " + err.Error())
		}
	}
}

func updateURLTable(url string, destination string, override bool) (bool, string) {
	record, exists := urlTable[url]

	if !exists || (exists && override) {
		urlTable[url] = destination
		return true, "Added: (" + url + " -> " + destination + ")"
	} else if exists && !override {
		return false, "Already exists (" + url + " -> " + record + "), please use override"
	}

	writeURLTable()
}
