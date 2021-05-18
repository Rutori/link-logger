package api

import (
	"context"
	"encoding/json"
	"net/http"

	"link-logger/db"
)

func List(_ context.Context, _ *http.Request) (response []byte, err error) {
	links, err := db.GetAllLinks()
	if err != nil {
		return nil, err
	}

	return json.Marshal(links)
}
