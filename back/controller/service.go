package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"link-logger/back/controller/credentials"
	"link-logger/interfaces"
)

type Service struct {
	Handles      map[string]*handle
	RootPath     string
	Password     string
	Verification credentials.Auth
}

// NewService creates service
func NewService(root string, authType credentials.Auth) *Service {
	return &Service{
		Handles:      make(map[string]*handle),
		Verification: authType,
		RootPath:     fmt.Sprintf("/%s", root),
	}
}

func (s *Service) RegisterHandle(path string, method string, handler func(ctx context.Context, request *http.Request) (response *Response, err error)) {
	switch method {
	case http.MethodPost, http.MethodGet:

	default:
		log.Fatalf("bad handle %s", "path")
	}
	handleName := s.RootPath
	if path != "" {
		handleName = fmt.Sprintf("%s/%s", handleName, path)
	}

	s.Handles[handleName] = &handle{
		Path:   path,
		Method: method,
		Func:   handler,
	}
}

func (s *Service) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if s.Verification != nil {
		auth := req.Header.Get("authorization")
		if auth == "" {
			auth = req.Header.Get("Authorization")
		}
		if !s.Verification.Verify(auth) {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	handler, correct := s.Handles[req.RequestURI]
	if !correct {
		handler, correct = s.Handles[fmt.Sprintf("%s/*", s.RootPath)] // try to catch the common handler
		if !correct {
			resp.WriteHeader(http.StatusNotFound)
			return
		}
	}

	if handler.Method != req.Method {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response, err := handler.Func(context.Background(), req)
	if err != nil {
		apiEr, isAPIerr := err.(interfaces.APIError)
		if !isAPIerr {
			log.Printf("%+v\n", err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Error(resp, apiEr.Error(), apiEr.Code())
		return
	}

	if response != nil {
		for k, v := range response.Headers {
			resp.Header().Set(k, v)
		}

		_, err = resp.Write(response.Body)
		if err != nil {
			log.Printf("%+v\n", err)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return
}
