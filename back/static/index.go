package static

import (
	"context"
	"io/ioutil"
	"net/http"

	"link-logger/back/controller"
)

func Index(_ context.Context, _ *http.Request) (response *controller.Response, err error) {
	body, err := ioutil.ReadFile("assets/html/index.html")
	response = &controller.Response{
		Body: body,
		Headers: map[string]string{
			"Content-Type": "text/html;charset=utf-8",
		},
	}

	return
}
