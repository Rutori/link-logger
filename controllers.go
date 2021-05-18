package main

import (
	"net/http"

	"link-logger/back/api"
	"link-logger/back/controller"
	"link-logger/back/static"
)

func controllers() http.Handler {

	apiController := controller.NewService("api", controller.ByPassword("not_a_real_password_duh"))
	apiController.RegisterHandle("add", http.MethodPost, api.Add)
	apiController.RegisterHandle("list", http.MethodGet, api.List)

	staticController := controller.NewService("", nil)
	staticController.RegisterHandle("", http.MethodGet, static.Index)

	backend := controller.CreateBackendHandler()
	backend.BindService(apiController)
	backend.BindService(staticController)

	return backend
}
