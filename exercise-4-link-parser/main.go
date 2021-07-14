package main

import "flag"

type Flags struct{
	htmlFilename string
}

func setupFlags() Flags {
	var htmlFilename string
	flag.StringVar(&htmlFilename, "htmlfile", "example1.html", "The HTML file you wish to extract the links from")
	flag.Parse()

	return Flags{
		htmlFilename,
	}
}

func main() {
	flags := setupFlags()
	links := GetLinksFromHtml(parseHtmlFile(flags.htmlFilename))

	printLinks(links)
}
