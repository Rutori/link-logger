package parser

import (
	"golang.org/x/net/html"
)

func title(document *html.Node) string {
	text, _ := traverse(document, isTitleElement)
	return text.FirstChild.Data
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}
