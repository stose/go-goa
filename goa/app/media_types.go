// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// An authorized response (default view)
//
// Identifier: application/vnd.authorized+json; view=default
type Authorized struct {
	// JWT token
	Token string `form:"token" json:"token" xml:"token"`
}

// Validate validates the Authorized media type instance.
func (mt *Authorized) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// A company information (default view)
//
// Identifier: application/vnd.company+json; view=default
type Company struct {
	Address string `form:"address" json:"address" xml:"address"`
	Country string `form:"country" json:"country" xml:"country"`
	// API href of company
	Href string `form:"href" json:"href" xml:"href"`
	// ID of company
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Company media type instance.
func (mt *Company) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}
	return
}

// A company information (link view)
//
// Identifier: application/vnd.company+json; view=link
type CompanyLink struct {
	// API href of company
	Href string `form:"href" json:"href" xml:"href"`
	// ID of company
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the CompanyLink media type instance.
func (mt *CompanyLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// A company information (tiny view)
//
// Identifier: application/vnd.company+json; view=tiny
type CompanyTiny struct {
	// API href of company
	Href string `form:"href" json:"href" xml:"href"`
	// ID of company
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the CompanyTiny media type instance.
func (mt *CompanyTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// A user information (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	Email string `form:"email" json:"email" xml:"email"`
	// User ID
	ID       *int   `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.UserName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	return
}

// A user information (id view)
//
// Identifier: application/vnd.user+json; view=id
type UserID struct {
	// User ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the UserID media type instance.
func (mt *UserID) Validate() (err error) {
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (id view)
//
// Identifier: application/vnd.user+json; type=collection; view=id
type UserIDCollection []*UserID

// Validate validates the UserIDCollection media type instance.
func (mt UserIDCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A user who belongs to which companies (default view)
//
// Identifier: application/vnd.usercomany+json; view=default
type Usercomany struct {
	// ID of user id
	CompanyID int `form:"company_id" json:"company_id" xml:"company_id"`
	// API href of bottle
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user company
	ID int `form:"id" json:"id" xml:"id"`
	// ID of user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Usercomany media type instance.
func (mt *Usercomany) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}
