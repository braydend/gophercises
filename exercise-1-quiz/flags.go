package main

import "flag"

type Flags struct {
	filename string
	timeLimit int
}

func setUpFlags() Flags {
	var filename string
	var timeLimit int
	flag.StringVar(&filename, "filename", "addition.csv", "A CSV to use for a quiz in the format (question,answer)")
	flag.IntVar(&timeLimit, "timeLimit", 30, "The time limit in seconds that the quiz should run for")

	flag.Parse()

	return Flags{filename,timeLimit}
}