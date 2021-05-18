package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"link-logger/back/controller"
	"link-logger/db"
)

func List(_ context.Context, _ *http.Request) (response *controller.Response, err error) {
	links, err := db.GetAllLinks()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := json.Marshal(links)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response = &controller.Response{
		Body: body,
	}

	return
}
