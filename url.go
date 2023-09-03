package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var urlTable map[string]string

const URLTABLEFLIE string = "/app/data/url_table.json"

func loadURLTable() bool {
	jsonFile, err := os.Open(URLTABLEFLIE)

	if err != nil {
		fmt.Println("Error while opening url_table.json")
		return false
	}

	defer jsonFile.Close()

	jsonBytes, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Cannot read url_table.json")
		return false
	}

	json.Unmarshal(jsonBytes, &urlTable)
	return true
}
