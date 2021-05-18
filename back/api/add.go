package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/mattn/go-sqlite3"
	"link-logger/db"
	"link-logger/db/models"
	"link-logger/parser"
)

type Exists struct{}

func (Exists) Code() int {
	return http.StatusBadRequest
}

func (Exists) Error() string {
	return "already exists"
}

func Add(_ context.Context, data []byte) (response []byte, err error) {
	pageURL, err := url.Parse(string(data))
	if err != nil {
		return nil, err
	}

	meta, err := parser.GetMeta(pageURL)
	if err != nil {
		return nil, err
	}

	res := db.Get().Create(&models.Links{
		Title:   meta.GetTitle(),
		Address: pageURL.String(),
		Preview: meta.GetPreview(),
	})

	if res.Error != nil {
		err, ok := res.Error.(sqlite3.Error)
		if !ok {
			return nil, res.Error
		}
		if err.ExtendedCode == sqlite3.ErrConstraintUnique {
			return nil, Exists{}
		}
	}

	return nil, nil
}
