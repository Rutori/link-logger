package static

import (
	"context"
	"io/ioutil"
	"net/http"
)

type FileNotFound struct{}

func (FileNotFound) Code() int {
	return http.StatusNotFound
}

func (FileNotFound) Error() string {
	return "asset is missing"
}

func Files(_ context.Context, req *http.Request) (response []byte, err error) {

	return ioutil.ReadFile(req.URL.Path)
}
