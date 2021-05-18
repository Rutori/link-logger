package parser

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

func fetchBody(url *url.URL) (io.Reader, error) {
	data, err := (&http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: 30 * time.Second,
	}).Get(url.String())

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if data.StatusCode != http.StatusOK {
		return nil, errors.New("Bad URL")
	}

	return data.Body, nil
}
