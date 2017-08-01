// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application User Types
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// commonResponse user type.
type commonResponse struct {
	// Date of creation
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Date of update
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the commonResponse type instance.
func (ut *commonResponse) Validate() (err error) {
	if ut.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.CreatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.created_at`, *ut.CreatedAt, goa.FormatDateTime, err2))
		}
	}
	if ut.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.UpdatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.updated_at`, *ut.UpdatedAt, goa.FormatDateTime, err2))
		}
	}
	return
}

// Publicize creates CommonResponse from commonResponse
func (ut *commonResponse) Publicize() *CommonResponse {
	var pub CommonResponse
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	return &pub
}

// CommonResponse user type.
type CommonResponse struct {
	// Date of creation
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Date of update
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the CommonResponse type instance.
func (ut *CommonResponse) Validate() (err error) {
	if ut.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.CreatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.created_at`, *ut.CreatedAt, goa.FormatDateTime, err2))
		}
	}
	if ut.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.UpdatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.updated_at`, *ut.UpdatedAt, goa.FormatDateTime, err2))
		}
	}
	return
}

// companyPayload user type.
type companyPayload struct {
	// Address of company
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country of HQ
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Name of company
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the companyPayload type instance.
func (ut *companyPayload) Validate() (err error) {
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) < 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 40, true))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) < 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 40, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 40, true))
		}
	}
	return
}

// Publicize creates CompanyPayload from companyPayload
func (ut *companyPayload) Publicize() *CompanyPayload {
	var pub CompanyPayload
	if ut.Address != nil {
		pub.Address = ut.Address
	}
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// CompanyPayload user type.
type CompanyPayload struct {
	// Address of company
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country of HQ
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Name of company
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the CompanyPayload type instance.
func (ut *CompanyPayload) Validate() (err error) {
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) < 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 40, true))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) < 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 40, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 40, true))
		}
	}
	return
}

// loginPayload user type.
type loginPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, true))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	return &pub
}

// LoginPayload user type.
type LoginPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, true))
		}
	}
	return
}

// userPayload user type.
type userPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// First name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, true))
		}
	}
	if ut.UserName != nil {
		if utf8.RuneCountInString(*ut.UserName) < 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user_name`, *ut.UserName, utf8.RuneCountInString(*ut.UserName), 20, true))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.UserName != nil {
		pub.UserName = ut.UserName
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// First name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, true))
		}
	}
	if ut.UserName != nil {
		if utf8.RuneCountInString(*ut.UserName) < 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user_name`, *ut.UserName, utf8.RuneCountInString(*ut.UserName), 20, true))
		}
	}
	return
}
