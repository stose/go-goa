package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyTechController implements the hy_tech resource.
type HyTechController struct {
	*goa.Controller
}

// NewHyTechController creates a hy_tech controller.
func NewHyTechController(service *goa.Service) *HyTechController {
	return &HyTechController{Controller: service.NewController("HyTechController")}
}

// CreateTech runs the CreateTech action.
func (c *HyTechController) CreateTech(ctx *app.CreateTechHyTechContext) error {
	// HyTechController_CreateTech: start_implement

	// Put your logic here

	res := &app.Tech{}
	return ctx.OK(res)
	// HyTechController_CreateTech: end_implement
}

// DeleteTech runs the DeleteTech action.
func (c *HyTechController) DeleteTech(ctx *app.DeleteTechHyTechContext) error {
	// HyTechController_DeleteTech: start_implement

	// Put your logic here

	res := &app.Tech{}
	return ctx.OK(res)
	// HyTechController_DeleteTech: end_implement
}

// GetTech runs the GetTech action.
func (c *HyTechController) GetTech(ctx *app.GetTechHyTechContext) error {
	// HyTechController_GetTech: start_implement

	// Put your logic here

	res := &app.Tech{}
	return ctx.OK(res)
	// HyTechController_GetTech: end_implement
}

// TechList runs the TechList action.
func (c *HyTechController) TechList(ctx *app.TechListHyTechContext) error {
	// HyTechController_TechList: start_implement

	// Put your logic here

	res := app.TechCollection{}
	return ctx.OK(res)
	// HyTechController_TechList: end_implement
}

// UpdateTech runs the UpdateTech action.
func (c *HyTechController) UpdateTech(ctx *app.UpdateTechHyTechContext) error {
	// HyTechController_UpdateTech: start_implement

	// Put your logic here

	res := &app.Tech{}
	return ctx.OK(res)
	// HyTechController_UpdateTech: end_implement
}
