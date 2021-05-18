package parser

import (
	"net/url"

	"golang.org/x/net/html"
)

func preview(document *html.Node) *url.URL {
	image, exists := traverse(document, isPreview)
	if !exists {
		return nil
	}

	var imageAddr *url.URL
	var err error
	for _, attr := range image.Attr {
		if attr.Key == "content" {
			imageAddr, err = url.Parse(attr.Val)
			if err != nil {
				return nil
			}
		}
	}

	return imageAddr
}

func isPreview(node *html.Node) bool {
	if node.Type != html.ElementNode || node.Data != "meta" {
		return false
	}

	for _, attr := range node.Attr {
		if attr.Key == "property" && attr.Val == "og:image" {
			return true
		}
	}

	return false
}
