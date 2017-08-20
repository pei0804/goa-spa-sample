package controller

import (
	"github.com/goadesign/goa"
	"github.com/pei0804/goa-spa-sample/app"
)

// HomeController implements the home resource.
type HomeController struct {
	*goa.Controller
}

// NewHomeController creates a home controller.
func NewHomeController(service *goa.Service) *HomeController {
	return &HomeController{Controller: service.NewController("HomeController")}
}

// Home runs the home action.
func (c *HomeController) Home(ctx *app.HomeHomeContext) error {
	// HomeController_Home: start_implement

	// Put your logic here

	// HomeController_Home: end_implement
	res := make(map[string]string, 1)
	res["message"] = "Hello World!"
	return ctx.OK(res)
}
