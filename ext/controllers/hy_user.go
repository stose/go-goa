// Package controllers is based on goa/hy_user.go generated by goagen automatically
package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	c "github.com/hiromaily/go-goa/ext/context"
	m "github.com/hiromaily/go-goa/ext/models"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyUserController implements something.
type HyUserController struct {
	*goa.Controller
	ctx *c.Ctx //Added
}

// NewHyUserController creates a HyUser controller.
func NewHyUserController(service *goa.Service, ctx *c.Ctx) *HyUserController {
	return &HyUserController{
		Controller: service.NewController("HyUserController"),
		ctx:        ctx, //Added
	}
}

// UserList runs the UserList action.
func (c *HyUserController) UserList(ctx *app.UserListHyUserContext) error {
	fmt.Println("[hy_user][UserList]")

	//type User struct {
	//	Email     string `form:"email" json:"email" xml:"email"`
	//	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	//	// User ID
	//	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	//}
	var users []*app.User

	svc := &m.User{Db: c.ctx.Db}
	err := svc.UserList(&users)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return ctx.NoContent()
	}

	//type UserCollection []*User
	res := app.UserCollection(users)
	return ctx.OK(res)
}

// GetUser runs the GetUser action.
func (c *HyUserController) GetUser(ctx *app.GetUserHyUserContext) error {
	fmt.Println("[hy_user][GetUser]")

	user := &app.User{}

	svc := &m.User{Db: c.ctx.Db}
	err := svc.GetUser(ctx.UserID, user)
	if err != nil {
		return err
	}

	if user.ID == nil {
		//404
		return ctx.NotFound()
	}

	//res := &app.User{}
	return ctx.OK(user)
}

// CreateUser runs the CreateUser action.
func (c *HyUserController) CreateUser(ctx *app.CreateUserHyUserContext) error {
	fmt.Println("[hy_user][CreateUser]")

	svc := &m.User{Db: c.ctx.Db}
	userID, err := svc.InsertUser(ctx.Payload) //*CreateUserHyUserPayload
	if err != nil {
		return err
	}

	res := &app.UserID{ID: &userID}
	return ctx.OKId(res)
}

// UpdateUser runs the UpdateUser action.
func (c *HyUserController) UpdateUser(ctx *app.UpdateUserHyUserContext) error {
	fmt.Println("[hy_user][UpdateUser]")

	svc := &m.User{Db: c.ctx.Db}
	err := svc.UpdateUser(ctx.UserID, ctx.Payload)
	if err != nil {
		return err
	}

	res := &app.UserID{ID: &ctx.UserID}
	return ctx.OKId(res)
}

// DeleteUser runs the DeleteUser action.
func (c *HyUserController) DeleteUser(ctx *app.DeleteUserHyUserContext) error {
	fmt.Println("[hy_user][DeleteUser]")

	svc := &m.User{Db: c.ctx.Db}
	err := svc.DeleteUser(ctx.UserID)
	if err != nil {
		return err
	}

	res := &app.UserID{ID: &ctx.UserID}
	return ctx.OKId(res)
}
