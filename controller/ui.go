package controller

import (
	"github.com/goadesign/goa"
)

// UIController implements the ui resource.
type UIController struct {
	*goa.Controller
}

// NewUIController creates a ui controller.
func NewUIController(service *goa.Service) *UIController {
	return &UIController{Controller: service.NewController("UIController")}
}
