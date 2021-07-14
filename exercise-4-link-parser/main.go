package main

func main() {
	links := GetLinksFromHtml(parseHtmlFile("example1.html"))

	printLinks(links)
}
