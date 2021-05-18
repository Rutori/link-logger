package controller

import (
	"fmt"
	"net/http"
	"strings"
)

type Backend map[string]*Service

func CreateBackendHandler() Backend {
	return Backend{}
}

func (mc Backend) BindService(service *Service) {
	mc[service.RootPath] = service
}

func (mc Backend) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var root string
	explodedPath := strings.Split(req.URL.Path, "/")
	if len(explodedPath) > 1 {
		root = fmt.Sprintf("/%s", explodedPath[1])
	}
	responseService, exists := mc[root]
	if !exists {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	responseService.ServeHTTP(resp, req)
}
