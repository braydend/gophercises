package common

import (
	"io/ioutil"
	"os"
)

func readFile(file *os.File) []byte {
	data, err := ioutil.ReadAll(file)
	handleError(err)

	return data
}

func OpenFile(filename string) *os.File {
	data, err := os.Open(filename)
	handleError(err)

	return data
}

func ParseFile(filename string) []byte {
	return readFile(OpenFile(filename))
}
