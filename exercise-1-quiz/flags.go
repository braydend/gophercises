package main

import "flag"

type Flags struct {
	filename string
}

func setUpFlags() Flags {
	var filename string
	flag.StringVar(&filename, "filename", "addition.csv", "A CSV to use for a quiz in the format (question,answer)")

	flag.Parse()

	return Flags{
		filename: filename,
	}
}