package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var urlTable = make(map[string]string)

const URLTABLEFLIE string = "data/url.csv"

func loadURLTable() bool {
	file, err := os.Open(URLTABLEFLIE)
	if err != nil {
		fmt.Print("error while opening url.csv: " + err.Error())
		return false
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	data, err := reader.ReadAll()
	if err != nil {
		fmt.Print("error while reading url.csv: " + err.Error())
		return false
	}

	for _, row := range data {
		urlTable[row[0]] = row[1]
	}

	fmt.Print(urlTable)
	return true
}
