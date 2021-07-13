package main

import "flag"

type Flags struct {
	storyFile string
}

func setupFlags() Flags{
	var storyFile string
	flag.StringVar(&storyFile, "storyFile", "gopher.json", "JSON file containing the story data. See gopher.json for an example of the structure")

	flag.Parse()

	return Flags{storyFile}
}
