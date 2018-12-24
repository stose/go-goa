// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "api": hy_companybranch Resource Client
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateCompanyBranchHyCompanybranchPayload is the hy_companybranch CreateCompanyBranch action payload.
type CreateCompanyBranchHyCompanybranchPayload struct {
	// Company Address
	Address string `form:"address" json:"address" yaml:"address" xml:"address"`
	// Country ID
	CountryID int `form:"country_id" json:"country_id" yaml:"country_id" xml:"country_id"`
}

// CreateCompanyBranchHyCompanybranchPath computes a request path to the CreateCompanyBranch action of hy_companybranch.
func CreateCompanyBranchHyCompanybranchPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/company/branch/%s", param0)
}

// Create new company branch
func (c *Client) CreateCompanyBranchHyCompanybranch(ctx context.Context, path string, payload *CreateCompanyBranchHyCompanybranchPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateCompanyBranchHyCompanybranchRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCompanyBranchHyCompanybranchRequest create the request corresponding to the CreateCompanyBranch action endpoint of the hy_companybranch resource.
func (c *Client) NewCreateCompanyBranchHyCompanybranchRequest(ctx context.Context, path string, payload *CreateCompanyBranchHyCompanybranchPayload, contentType string) (*http.Request, error) {
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
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteCompanyBranchHyCompanybranchPath computes a request path to the DeleteCompanyBranch action of hy_companybranch.
func DeleteCompanyBranchHyCompanybranchPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/company/branch/%s", param0)
}

// Delete company branch
func (c *Client) DeleteCompanyBranchHyCompanybranch(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCompanyBranchHyCompanybranchRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCompanyBranchHyCompanybranchRequest create the request corresponding to the DeleteCompanyBranch action endpoint of the hy_companybranch resource.
func (c *Client) NewDeleteCompanyBranchHyCompanybranchRequest(ctx context.Context, path string) (*http.Request, error) {
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
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetCompanyBranchHyCompanybranchPath computes a request path to the GetCompanyBranch action of hy_companybranch.
func GetCompanyBranchHyCompanybranchPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/company/branch/%s", param0)
}

// Retrieve company branch with given id
func (c *Client) GetCompanyBranchHyCompanybranch(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetCompanyBranchHyCompanybranchRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCompanyBranchHyCompanybranchRequest create the request corresponding to the GetCompanyBranch action endpoint of the hy_companybranch resource.
func (c *Client) NewGetCompanyBranchHyCompanybranchRequest(ctx context.Context, path string) (*http.Request, error) {
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
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateCompanyBranchHyCompanybranchPayload is the hy_companybranch UpdateCompanyBranch action payload.
type UpdateCompanyBranchHyCompanybranchPayload struct {
	// Company Address
	Address string `form:"address" json:"address" yaml:"address" xml:"address"`
	// Country ID
	CountryID int `form:"country_id" json:"country_id" yaml:"country_id" xml:"country_id"`
}

// UpdateCompanyBranchHyCompanybranchPath computes a request path to the UpdateCompanyBranch action of hy_companybranch.
func UpdateCompanyBranchHyCompanybranchPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/company/branch/%s", param0)
}

// Change company branch properties
func (c *Client) UpdateCompanyBranchHyCompanybranch(ctx context.Context, path string, payload *UpdateCompanyBranchHyCompanybranchPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateCompanyBranchHyCompanybranchRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCompanyBranchHyCompanybranchRequest create the request corresponding to the UpdateCompanyBranch action endpoint of the hy_companybranch resource.
func (c *Client) NewUpdateCompanyBranchHyCompanybranchRequest(ctx context.Context, path string, payload *UpdateCompanyBranchHyCompanybranchPayload, contentType string) (*http.Request, error) {
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
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
