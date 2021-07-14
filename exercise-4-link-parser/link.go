package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func printLinks(links []Link) {
	fmt.Printf("Found %d link(s) in the document:\n", len(links))
	for _, link := range links {
		fmt.Printf("(%s) with the label: %s\n", link.Href, link.Text)
	}
}

func GetLinksFromHtml(doc *html.Node) (links []Link) {
	appendToList := func (node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			links = append(links, getLinkDataFromNode(node))
		}
	}

	traverseHtmlTree(doc, appendToList)

	return links
}

func getLinkDataFromNode(node *html.Node) Link {
	var href string
	var text string

	getTextForLink := func (n *html.Node) {
		if n.Type == html.TextNode {
			text += n.Data
		}
	}

	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			href = attribute.Val
		}
	}

	traverseHtmlTree(node, getTextForLink)

	return Link{href, strings.TrimSpace(text)}
}
