//go:generate goagen bootstrap -d github.com/pei0804/goa-spa-sample/design

package server

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/pei0804/goa-spa-sample/app"
	"github.com/pei0804/goa-spa-sample/controller"
	"github.com/pei0804/goa-spa-sample/front"
)

func init() {
	// Create service
	service := goa.New("spa")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "front" controller
	c := controller.NewFrontController(service)
	c.FileSystem = func(dir string) http.FileSystem {
		return &assetfs.AssetFS{
			Asset:     front.Asset,
			AssetDir:  front.AssetDir,
			AssetInfo: front.AssetInfo,
			Prefix:    dir,
		}
	}
	app.MountFrontController(service, c)
	// Mount "home" controller
	c2 := controller.NewHomeController(service)
	app.MountHomeController(service, c2)
	// Mount "schema" controller
	c3 := controller.NewSchemaController(service)
	app.MountSchemaController(service, c3)
	// Mount "swagger" controller
	c4 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c4)

	// Start service
	http.HandleFunc("/", service.Mux.ServeHTTP)

}
