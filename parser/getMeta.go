package parser

import (
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"link-logger/interfaces"
)

// GetMeta returns metadata
func GetMeta(url *url.URL) (interfaces.Metadata, error) {
	bodyReader, err := fetchBody(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	htmlData, err := html.Parse(bodyReader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &meta{
		Title:       html.EscapeString(title(htmlData)),
		PreviewURL:  preview(htmlData),
		Description: description(htmlData),
	}, nil
}
