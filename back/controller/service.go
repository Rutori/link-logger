package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"link-logger/interfaces"
)

type Service struct {
	Handles      map[string]*handle
	RootPath     string
	Password     string
	Verification Auth
}

// NewService creates service
func NewService(root string, authType Auth) *Service {
	return &Service{
		Handles:      make(map[string]*handle),
		Verification: authType,
		RootPath:     fmt.Sprintf("/%s", root),
	}
}

func (s *Service) RegisterHandle(path string, method string, handler func(ctx context.Context, request []byte) (response []byte, err error)) {
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
	if s.Verification != nil && s.Verification.Verify(req.Header.Get("Authorization")) == false {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	handler, correct := s.Handles[req.RequestURI]
	if !correct {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	if handler.Method != req.Method {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := handler.Func(context.Background(), body)
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

	_, err = resp.Write(response)
	if err != nil {
		log.Printf("%+v\n", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
