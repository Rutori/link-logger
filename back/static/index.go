package static

import (
	"context"
	"io/ioutil"
)

func Index(_ context.Context, _ []byte) (response []byte, err error) {
	response, err = ioutil.ReadFile("assets/html/index.html")
	return
}
