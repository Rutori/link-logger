package api

import "net/http"

type Unreadable struct{}

func (Unreadable) Code() int {
	return http.StatusBadRequest
}

func (Unreadable) Error() string {
	return "bad request format"
}
