package static

import (
	"context"
	"io/ioutil"
	"net/http"
)

func Index(_ context.Context, _ *http.Request) (response []byte, err error) {
	response, err = ioutil.ReadFile("assets/html/index.html")
	return
}
