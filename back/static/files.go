package static

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"link-logger/back/controller"
)

type FileNotFound struct{}

func (FileNotFound) Code() int {
	return http.StatusNotFound
}

func (FileNotFound) Error() string {
	return "asset is missing"
}

func Files(_ context.Context, req *http.Request) (response *controller.Response, err error) {
	body, err := ioutil.ReadFile(strings.TrimLeft(req.URL.Path, "/"))
	if errors.Is(err, &os.PathError{}) {
		return nil, FileNotFound{}
	}

	contentType, allowed := fetchContentType(req.URL.Path)
	if !allowed {
		return nil, FileNotFound{}
	}

	response = &controller.Response{
		Body: body,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
	}

	return
}

func fetchContentType(fileName string) (contentType string, exists bool) {
	contentType, exists = mimeMap[path.Ext(fileName)]

	return fmt.Sprintf("%s; charset=utf-8", contentType), exists
}
