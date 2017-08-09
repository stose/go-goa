// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
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

// DecodeAuthorized decodes the Authorized instance encoded in resp body.
func (c *Client) DecodeAuthorized(resp *http.Response) (*Authorized, error) {
	var decoded Authorized
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A company information (default view)
//
// Identifier: application/vnd.company+json; view=default
type Company struct {
	Address string `form:"address" json:"address" xml:"address"`
	// ID
	CompanyID   *int    `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" xml:"country_name,omitempty"`
	HqFlg       *string `form:"hq_flg,omitempty" json:"hq_flg,omitempty" xml:"hq_flg,omitempty"`
	// ID
	ID   *int   `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name string `form:"name" json:"name" xml:"name"`
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
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
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
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
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
	CompanyID *int   `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	Name      string `form:"name" json:"name" xml:"name"`
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

// DecodeCompany decodes the Company instance encoded in resp body.
func (c *Client) DecodeCompany(resp *http.Response) (*Company, error) {
	var decoded Company
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCompanyDetailid decodes the CompanyDetailid instance encoded in resp body.
func (c *Client) DecodeCompanyDetailid(resp *http.Response) (*CompanyDetailid, error) {
	var decoded CompanyDetailid
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCompanyID decodes the CompanyID instance encoded in resp body.
func (c *Client) DecodeCompanyID(resp *http.Response) (*CompanyID, error) {
	var decoded CompanyID
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCompanyIdname decodes the CompanyIdname instance encoded in resp body.
func (c *Client) DecodeCompanyIdname(resp *http.Response) (*CompanyIdname, error) {
	var decoded CompanyIdname
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeCompanyCollection decodes the CompanyCollection instance encoded in resp body.
func (c *Client) DecodeCompanyCollection(resp *http.Response) (CompanyCollection, error) {
	var decoded CompanyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeCompanyDetailidCollection decodes the CompanyDetailidCollection instance encoded in resp body.
func (c *Client) DecodeCompanyDetailidCollection(resp *http.Response) (CompanyDetailidCollection, error) {
	var decoded CompanyDetailidCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeCompanyIDCollection decodes the CompanyIDCollection instance encoded in resp body.
func (c *Client) DecodeCompanyIDCollection(resp *http.Response) (CompanyIDCollection, error) {
	var decoded CompanyIDCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeCompanyIdnameCollection decodes the CompanyIdnameCollection instance encoded in resp body.
func (c *Client) DecodeCompanyIdnameCollection(resp *http.Response) (CompanyIdnameCollection, error) {
	var decoded CompanyIdnameCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A user information (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	Email string `form:"email" json:"email" xml:"email"`
	// ID
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
	// ID
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

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserID decodes the UserID instance encoded in resp body.
func (c *Client) DecodeUserID(resp *http.Response) (*UserID, error) {
	var decoded UserID
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeUserIDCollection decodes the UserIDCollection instance encoded in resp body.
func (c *Client) DecodeUserIDCollection(resp *http.Response) (UserIDCollection, error) {
	var decoded UserIDCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
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

// DecodeUsercomany decodes the Usercomany instance encoded in resp body.
func (c *Client) DecodeUsercomany(resp *http.Response) (*Usercomany, error) {
	var decoded Usercomany
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A user information (default view)
//
// Identifier: application/vnd.usertech+json; view=default
type Usertech struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Tech name
	TechName string `form:"tech_name" json:"tech_name" xml:"tech_name"`
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
	if utf8.RuneCountInString(mt.TechName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 2, true))
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
	TechName string `form:"tech_name" json:"tech_name" xml:"tech_name"`
}

// Validate validates the UsertechTech media type instance.
func (mt *UsertechTech) Validate() (err error) {
	if mt.TechName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tech_name"))
	}
	if utf8.RuneCountInString(mt.TechName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 2, true))
	}
	if utf8.RuneCountInString(mt.TechName) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.tech_name`, mt.TechName, utf8.RuneCountInString(mt.TechName), 40, false))
	}
	return
}

// DecodeUsertech decodes the Usertech instance encoded in resp body.
func (c *Client) DecodeUsertech(resp *http.Response) (*Usertech, error) {
	var decoded Usertech
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUsertechTech decodes the UsertechTech instance encoded in resp body.
func (c *Client) DecodeUsertechTech(resp *http.Response) (*UsertechTech, error) {
	var decoded UsertechTech
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeUsertechCollection decodes the UsertechCollection instance encoded in resp body.
func (c *Client) DecodeUsertechCollection(resp *http.Response) (UsertechCollection, error) {
	var decoded UsertechCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeUsertechTechCollection decodes the UsertechTechCollection instance encoded in resp body.
func (c *Client) DecodeUsertechTechCollection(resp *http.Response) (UsertechTechCollection, error) {
	var decoded UsertechTechCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A user information (default view)
//
// Identifier: application/vnd.userworkhistory+json; view=default
type Userworkhistory struct {
	// Company name
	Company string `form:"company" json:"company" xml:"company"`
	// Country code
	Country string `form:"country" json:"country" xml:"country"`
	// job description
	Description *interface{} `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// used techs
	Techs *interface{} `form:"techs,omitempty" json:"techs,omitempty" xml:"techs,omitempty"`
	// worked period
	Term *string `form:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// Job Title
	Title string `form:"title" json:"title" xml:"title"`
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

// DecodeUserworkhistory decodes the Userworkhistory instance encoded in resp body.
func (c *Client) DecodeUserworkhistory(resp *http.Response) (*Userworkhistory, error) {
	var decoded Userworkhistory
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeUserworkhistoryCollection decodes the UserworkhistoryCollection instance encoded in resp body.
func (c *Client) DecodeUserworkhistoryCollection(resp *http.Response) (UserworkhistoryCollection, error) {
	var decoded UserworkhistoryCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
