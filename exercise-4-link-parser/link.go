package main

import (
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
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
			text = n.Data
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
