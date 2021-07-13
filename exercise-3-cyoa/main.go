package main

func main()  {
	flags := setupFlags()
	startServer(parseStories(parseFile(flags.storyFile)))
}
