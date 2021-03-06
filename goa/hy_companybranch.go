package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyCompanybranchController implements the hy_companybranch resource.
type HyCompanybranchController struct {
	*goa.Controller
}

// NewHyCompanybranchController creates a hy_companybranch controller.
func NewHyCompanybranchController(service *goa.Service) *HyCompanybranchController {
	return &HyCompanybranchController{Controller: service.NewController("HyCompanybranchController")}
}

// CreateCompanyBranch runs the CreateCompanyBranch action.
func (c *HyCompanybranchController) CreateCompanyBranch(ctx *app.CreateCompanyBranchHyCompanybranchContext) error {
	// HyCompanybranchController_CreateCompanyBranch: start_implement

	// Put your logic here

	res := &app.Company{}
	return ctx.OK(res)
	// HyCompanybranchController_CreateCompanyBranch: end_implement
}

// DeleteCompanyBranch runs the DeleteCompanyBranch action.
func (c *HyCompanybranchController) DeleteCompanyBranch(ctx *app.DeleteCompanyBranchHyCompanybranchContext) error {
	// HyCompanybranchController_DeleteCompanyBranch: start_implement

	// Put your logic here

	res := &app.Company{}
	return ctx.OK(res)
	// HyCompanybranchController_DeleteCompanyBranch: end_implement
}

// GetCompanyBranch runs the GetCompanyBranch action.
func (c *HyCompanybranchController) GetCompanyBranch(ctx *app.GetCompanyBranchHyCompanybranchContext) error {
	// HyCompanybranchController_GetCompanyBranch: start_implement

	// Put your logic here

	res := &app.Company{}
	return ctx.OK(res)
	// HyCompanybranchController_GetCompanyBranch: end_implement
}

// UpdateCompanyBranch runs the UpdateCompanyBranch action.
func (c *HyCompanybranchController) UpdateCompanyBranch(ctx *app.UpdateCompanyBranchHyCompanybranchContext) error {
	// HyCompanybranchController_UpdateCompanyBranch: start_implement

	// Put your logic here

	res := &app.Company{}
	return ctx.OK(res)
	// HyCompanybranchController_UpdateCompanyBranch: end_implement
}
