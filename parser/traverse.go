package parser

import "golang.org/x/net/html"

func traverse(n *html.Node, validator func(node *html.Node) bool) (*html.Node, bool) {
	if validator(n) {
		return n, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c, validator)
		if ok {
			return result, ok
		}
	}

	return nil, false
}
