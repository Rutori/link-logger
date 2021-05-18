package back

import (
	"net/http"

	"link-logger/back/api"
	"link-logger/back/controller"
	"link-logger/back/controller/credentials"
	"link-logger/back/static"
	"link-logger/config"
)

func Init() http.Handler {

	apiController := controller.NewService("api", credentials.ByPassword(config.Storage.AdminPassword))
	apiController.RegisterHandle("add", http.MethodPost, api.Add)
	apiController.RegisterHandle("list", http.MethodGet, api.List)

	pageController := controller.NewService("", nil)
	pageController.RegisterHandle("", http.MethodGet, static.Index)

	staticController := controller.NewService("assets", nil)
	staticController.RegisterHandle("*", http.MethodGet, static.Files)

	backend := controller.CreateBackendHandler()
	backend.BindService(apiController)
	backend.BindService(pageController)
	backend.BindService(staticController)

	return backend
}
