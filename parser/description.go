package parser

import (
	"golang.org/x/net/html"
)

func description(document *html.Node) string {
	desc, _ := traverse(document, isDescriptionElement)
	for _, attr := range desc.Attr {
		if attr.Key == "content" {
			return attr.Val
		}
	}

	return ""
}

func isDescriptionElement(node *html.Node) bool {
	if node.Type != html.ElementNode || node.Data != "meta" {
		return false
	}

	for _, attr := range node.Attr {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}

		if attr.Key == "property" && attr.Val == "og:description" {
			return true
		}
	}

	return false
}
