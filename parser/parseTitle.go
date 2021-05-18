package parser

import (
	"golang.org/x/net/html"
)

func parseTitle(document *html.Node) string {
	title, _ := traverse(document, isTitleElement)
	return title.FirstChild.Data
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}
