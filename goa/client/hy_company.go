// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": hy_company Resource Client
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CompanyListHyCompanyPath computes a request path to the CompanyList action of hy_company.
func CompanyListHyCompanyPath() string {

	return fmt.Sprintf("/api/company")
}

// List all companies
func (c *Client) CompanyListHyCompany(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCompanyListHyCompanyRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCompanyListHyCompanyRequest create the request corresponding to the CompanyList action endpoint of the hy_company resource.
func (c *Client) NewCompanyListHyCompanyRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// CreateCompanyHyCompanyPayload is the hy_company CreateCompany action payload.
type CreateCompanyHyCompanyPayload struct {
	// Address of company
	Address string `form:"address" json:"address" xml:"address"`
	// Company ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Country's ID
	CountryID string `form:"country_id" json:"country_id" xml:"country_id"`
	// Headquarters flg
	HqFlg *string `form:"hq_flg,omitempty" json:"hq_flg,omitempty" xml:"hq_flg,omitempty"`
	// Company Name
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateCompanyHyCompanyPath computes a request path to the CreateCompany action of hy_company.
func CreateCompanyHyCompanyPath() string {

	return fmt.Sprintf("/api/company")
}

// Create new company
func (c *Client) CreateCompanyHyCompany(ctx context.Context, path string, payload *CreateCompanyHyCompanyPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCompanyHyCompanyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCompanyHyCompanyRequest create the request corresponding to the CreateCompany action endpoint of the hy_company resource.
func (c *Client) NewCreateCompanyHyCompanyRequest(ctx context.Context, path string, payload *CreateCompanyHyCompanyPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/xml")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// DeleteCompanyHyCompanyPath computes a request path to the DeleteCompany action of hy_company.
func DeleteCompanyHyCompanyPath(companyID int) string {
	param0 := strconv.Itoa(companyID)

	return fmt.Sprintf("/api/company/%s", param0)
}

// Delete company
func (c *Client) DeleteCompanyHyCompany(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCompanyHyCompanyRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCompanyHyCompanyRequest create the request corresponding to the DeleteCompany action endpoint of the hy_company resource.
func (c *Client) NewDeleteCompanyHyCompanyRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// GetCompanyGroupHyCompanyPath computes a request path to the GetCompanyGroup action of hy_company.
func GetCompanyGroupHyCompanyPath(companyID int) string {
	param0 := strconv.Itoa(companyID)

	return fmt.Sprintf("/api/company/%s", param0)
}

// Retrieve company with given company_id
func (c *Client) GetCompanyGroupHyCompany(ctx context.Context, path string, hqFlg *string) (*http.Response, error) {
	req, err := c.NewGetCompanyGroupHyCompanyRequest(ctx, path, hqFlg)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCompanyGroupHyCompanyRequest create the request corresponding to the GetCompanyGroup action endpoint of the hy_company resource.
func (c *Client) NewGetCompanyGroupHyCompanyRequest(ctx context.Context, path string, hqFlg *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if hqFlg != nil {
		values.Set("hq_flg", *hqFlg)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// UpdateCompanyHyCompanyPayload is the hy_company UpdateCompany action payload.
type UpdateCompanyHyCompanyPayload struct {
	// Address of company
	Address string `form:"address" json:"address" xml:"address"`
	// Company ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Country's ID
	CountryID string `form:"country_id" json:"country_id" xml:"country_id"`
	// Headquarters flg
	HqFlg *string `form:"hq_flg,omitempty" json:"hq_flg,omitempty" xml:"hq_flg,omitempty"`
	// Company Name
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateCompanyHyCompanyPath computes a request path to the UpdateCompany action of hy_company.
func UpdateCompanyHyCompanyPath(companyID int) string {
	param0 := strconv.Itoa(companyID)

	return fmt.Sprintf("/api/company/%s", param0)
}

// Change company properties
func (c *Client) UpdateCompanyHyCompany(ctx context.Context, path string, payload *UpdateCompanyHyCompanyPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateCompanyHyCompanyRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCompanyHyCompanyRequest create the request corresponding to the UpdateCompany action endpoint of the hy_company resource.
func (c *Client) NewUpdateCompanyHyCompanyRequest(ctx context.Context, path string, payload *UpdateCompanyHyCompanyPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/xml")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
