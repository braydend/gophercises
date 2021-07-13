package main

func main()  {
	startServer(parseStories(parseFile("gopher.json")))
}
