package parser

import "net/url"

type meta struct {
	Title       string
	Description string
	PreviewURL  *url.URL
}

func (m meta) GetDescription() string {
	return m.Description
}

func (m meta) GetTitle() string {
	return m.Title
}

func (m meta) GetPreview() string {
	if m.PreviewURL == nil {
		return ""
	}
	return m.PreviewURL.String()
}
