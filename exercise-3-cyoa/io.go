package main

import (
	"io/ioutil"
	"log"
	"os"
)

func readFile(file *os.File) []byte {
	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return fileBytes
}

func parseFile(filename string) []byte {
	file := openFile(filename)
	return readFile(file)
}

func openFile(filename string) *os.File{
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
