package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

func openFile(filename string) *os.File {
	data, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func readCsv(filename string) [][]string {
	records, err := csv.NewReader(openFile(filename)).ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}

func readInput() []byte {
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	return input
}
