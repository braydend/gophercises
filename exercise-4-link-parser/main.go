package main

import "fmt"

func main() {
	links := GetLinksFromHtml(parseHtmlFile("example1.html"))

	fmt.Println(links)
}
