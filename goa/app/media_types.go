// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "api": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// An authorized response (default view)
//
// Identifier: application/vnd.authorized+json; view=default
type Authorized struct {
	// ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// JWT token
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
}

// Validate validates the Authorized media type instance.
func (mt *Authorized) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	return
}

// A company information (default view)
//
// Identifier: application/vnd.company+json; view=default
type Company struct {
	Address string `form:"address" json:"address" yaml:"address" xml:"address"`
	// ID
	CompanyID   *int    `form:"company_id,omitempty" json:"company_id,omitempty" yaml:"company_id,omitempty" xml:"company_id,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" yaml:"country_name,omitempty" xml:"country_name,omitempty"`
	HqFlg       *string `form:"hq_flg,omitempty" json:"hq_flg,omitempty" yaml:"hq_flg,omitempty" xml:"hq_flg,omitempty"`
	// ID
	ID   *int   `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the Company media type instance.
func (mt *Company) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}
	if mt.CompanyID != nil {
		if *mt.CompanyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.company_id`, *mt.CompanyID, 1, true))
		}
	}
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	return
}

// A company information (detailid view)
//
// Identifier: application/vnd.company+json; view=detailid
type CompanyDetailid struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the CompanyDetailid media type instance.
func (mt *CompanyDetailid) Validate() (err error) {
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	return
}

// A company information (id view)
//
// Identifier: application/vnd.company+json; view=id
type CompanyID struct {
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" yaml:"company_id,omitempty" xml:"company_id,omitempty"`
}

// Validate validates the CompanyID media type instance.
func (mt *CompanyID) Validate() (err error) {
	if mt.CompanyID != nil {
		if *mt.CompanyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.company_id`, *mt.CompanyID, 1, true))
		}
	}
	return
}

// A company information (idname view)
//
// Identifier: application/vnd.company+json; view=idname
type CompanyIdname struct {
	// ID
	CompanyID *int   `form:"company_id,omitempty" json:"company_id,omitempty" yaml:"company_id,omitempty" xml:"company_id,omitempty"`
	Name      string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the CompanyIdname media type instance.
func (mt *CompanyIdname) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.CompanyID != nil {
		if *mt.CompanyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.company_id`, *mt.CompanyID, 1, true))
		}
	}
	return
}

// CompanyCollection is the media type for an array of Company (default view)
//
// Identifier: application/vnd.company+json; type=collection; view=default
type CompanyCollection []*Company

// Validate validates the CompanyCollection media type instance.
func (mt CompanyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// CompanyCollection is the media type for an array of Company (detailid view)
//
// Identifier: application/vnd.company+json; type=collection; view=detailid
type CompanyDetailidCollection []*CompanyDetailid

// Validate validates the CompanyDetailidCollection media type instance.
func (mt CompanyDetailidCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// CompanyCollection is the media type for an array of Company (id view)
//
// Identifier: application/vnd.company+json; type=collection; view=id
type CompanyIDCollection []*CompanyID

// Validate validates the CompanyIDCollection media type instance.
func (mt CompanyIDCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// CompanyCollection is the media type for an array of Company (idname view)
//
// Identifier: application/vnd.company+json; type=collection; view=idname
type CompanyIdnameCollection []*CompanyIdname

// Validate validates the CompanyIdnameCollection media type instance.
func (mt CompanyIdnameCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A tech information (default view)
//
// Identifier: application/vnd.tech+json; view=default
type Tech struct {
	// ID
	ID   *int   `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the Tech media type instance.
func (mt *Tech) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	return
}

// A tech information (id view)
//
// Identifier: application/vnd.tech+json; view=id
type TechID struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the TechID media type instance.
func (mt *TechID) Validate() (err error) {
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	return
}

// TechCollection is the media type for an array of Tech (default view)
//
// Identifier: application/vnd.tech+json; type=collection; view=default
type TechCollection []*Tech

// Validate validates the TechCollection media type instance.
func (mt TechCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// TechCollection is the media type for an array of Tech (id view)
//
// Identifier: application/vnd.tech+json; type=collection; view=id
type TechIDCollection []*TechID

// Validate validates the TechIDCollection media type instance.
func (mt TechIDCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A user information (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	// ID
	ID       *int   `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	UserName string `form:"user_name" json:"user_name" yaml:"user_name" xml:"user_name"`
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
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
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
	CompanyID int `form:"company_id" json:"company_id" yaml:"company_id" xml:"company_id"`
	// API href of bottle
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// ID of user company
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// ID of user id
	UserID int `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the Usercomany media type instance.
func (mt *Usercomany) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}

// A user information (default view)
//
// Identifier: application/vnd.usertech+json; view=default
type Usertech struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Tech name
	TechName string `form:"tech_name" json:"tech_name" yaml:"tech_name" xml:"tech_name"`
}

// Validate validates the Usertech media type instance.
func (mt *Usertech) Validate() (err error) {
	if mt.TechName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tech_name"))
	}
	if mt.ID != nil {
		if *mt.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *mt.ID, 1, true))
		}
	}
	if utf8.RuneCountInString(mt.TechName) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 1, true))
	}
	if utf8.RuneCountInString(mt.TechName) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 40, false))
	}
	return
}

// A user information (tech view)
//
// Identifier: application/vnd.usertech+json; view=tech
type UsertechTech struct {
	// Tech name
	TechName string `form:"tech_name" json:"tech_name" yaml:"tech_name" xml:"tech_name"`
}

// Validate validates the UsertechTech media type instance.
func (mt *UsertechTech) Validate() (err error) {
	if mt.TechName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tech_name"))
	}
	if utf8.RuneCountInString(mt.TechName) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 1, true))
	}
	if utf8.RuneCountInString(mt.TechName) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 40, false))
	}
	return
}

// UsertechCollection is the media type for an array of Usertech (default view)
//
// Identifier: application/vnd.usertech+json; type=collection; view=default
type UsertechCollection []*Usertech

// Validate validates the UsertechCollection media type instance.
func (mt UsertechCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UsertechCollection is the media type for an array of Usertech (tech view)
//
// Identifier: application/vnd.usertech+json; type=collection; view=tech
type UsertechTechCollection []*UsertechTech

// Validate validates the UsertechTechCollection media type instance.
func (mt UsertechTechCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A user information (default view)
//
// Identifier: application/vnd.userworkhistory+json; view=default
type Userworkhistory struct {
	// Company name
	Company string `form:"company" json:"company" yaml:"company" xml:"company"`
	// Country code
	Country string `form:"country" json:"country" yaml:"country" xml:"country"`
	// job description
	Description interface{} `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// used techs
	Techs interface{} `form:"techs,omitempty" json:"techs,omitempty" yaml:"techs,omitempty" xml:"techs,omitempty"`
	// worked period
	Term *string `form:"term,omitempty" json:"term,omitempty" yaml:"term,omitempty" xml:"term,omitempty"`
	// Job Title
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// Validate validates the Userworkhistory media type instance.
func (mt *Userworkhistory) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Company == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "company"))
	}
	if mt.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}
	if utf8.RuneCountInString(mt.Company) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.company`, mt.Company, utf8.RuneCountInString(mt.Company), 2, true))
	}
	if utf8.RuneCountInString(mt.Company) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.company`, mt.Company, utf8.RuneCountInString(mt.Company), 40, false))
	}
	if utf8.RuneCountInString(mt.Country) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, mt.Country, utf8.RuneCountInString(mt.Country), 2, true))
	}
	if utf8.RuneCountInString(mt.Country) > 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, mt.Country, utf8.RuneCountInString(mt.Country), 2, false))
	}
	if mt.Term != nil {
		if utf8.RuneCountInString(*mt.Term) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.term`, *mt.Term, utf8.RuneCountInString(*mt.Term), 10, true))
		}
	}
	if mt.Term != nil {
		if utf8.RuneCountInString(*mt.Term) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.term`, *mt.Term, utf8.RuneCountInString(*mt.Term), 20, false))
		}
	}
	if utf8.RuneCountInString(mt.Title) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 2, true))
	}
	if utf8.RuneCountInString(mt.Title) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 40, false))
	}
	return
}

// UserworkhistoryCollection is the media type for an array of Userworkhistory (default view)
//
// Identifier: application/vnd.userworkhistory+json; type=collection; view=default
type UserworkhistoryCollection []*Userworkhistory

// Validate validates the UserworkhistoryCollection media type instance.
func (mt UserworkhistoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
