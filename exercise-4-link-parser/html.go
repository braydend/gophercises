package main

import (
	"github.com/braydend/gophercises/common"
	"golang.org/x/net/html"
	"log"
)

func parseHtmlFile(filename string) *html.Node {
	doc, err := html.Parse(common.OpenFile(filename))

	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func traverseHtmlTree(node *html.Node, callback func(n *html.Node)) {
	callback(node)

	// Recursively traverse the tree
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverseHtmlTree(child, callback)
	}
}
