package controller

import (
	"fmt"
	"net/http"
	"strings"
)

type Collection map[string]*Service

func CreateBackendHandler() Collection {
	return Collection{}
}

func (mc Collection) BindService(service *Service) {
	mc[service.RootPath] = service
}

func (mc Collection) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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
