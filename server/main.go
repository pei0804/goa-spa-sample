//go:generate goagen bootstrap -d github.com/pei0804/goa-spa-sample/design

package main

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/pei0804/goa-spa-sample/app"
	"github.com/pei0804/goa-spa-sample/controller"
)

func init() {
	// Create service
	service := goa.New("files")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "home" controller
	c := controller.NewHomeController(service)
	app.MountHomeController(service, c)
	// Mount "schema" controller
	c2 := controller.NewSchemaController(service)
	app.MountSchemaController(service, c2)
	// Mount "swagger" controller
	c3 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "ui" controller
	c4 := controller.NewUIController(service)
	app.MountUIController(service, c4)

	// Start service
	http.HandleFunc("/", service.Mux.ServeHTTP)

}
