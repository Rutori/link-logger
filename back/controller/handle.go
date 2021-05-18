package controller

import (
	"context"
	"net/http"
)

type handle struct {
	Path   string
	Method string
	Func   func(ctx context.Context, request *http.Request) (response *Response, err error)
}
