package api

import (
	"context"
	"encoding/json"

	"link-logger/db"
)

func List(_ context.Context, _ []byte) (response []byte, err error) {
	links, err := db.GetAllLinks()
	if err != nil {
		return nil, err
	}

	return json.Marshal(links)
}
