package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"
)

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func getCsvReader(filename string) *csv.Reader {
	file := readFile(filename)

	return csv.NewReader(strings.NewReader(string (file)))
}

func readCsv(reader *csv.Reader) [][]string {
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}
