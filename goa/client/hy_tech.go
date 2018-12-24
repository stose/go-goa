// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "api": hy_tech Resource Client
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

// CreateTechHyTechPayload is the hy_tech CreateTech action payload.
type CreateTechHyTechPayload struct {
	// Tech name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// CreateTechHyTechPath computes a request path to the CreateTech action of hy_tech.
func CreateTechHyTechPath() string {

	return fmt.Sprintf("/api/tech")
}

// Create new tech
func (c *Client) CreateTechHyTech(ctx context.Context, path string, payload *CreateTechHyTechPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateTechHyTechRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTechHyTechRequest create the request corresponding to the CreateTech action endpoint of the hy_tech resource.
func (c *Client) NewCreateTechHyTechRequest(ctx context.Context, path string, payload *CreateTechHyTechPayload, contentType string) (*http.Request, error) {
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

// DeleteTechHyTechPath computes a request path to the DeleteTech action of hy_tech.
func DeleteTechHyTechPath(techID int) string {
	param0 := strconv.Itoa(techID)

	return fmt.Sprintf("/api/tech/%s", param0)
}

// Delete tech
func (c *Client) DeleteTechHyTech(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteTechHyTechRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteTechHyTechRequest create the request corresponding to the DeleteTech action endpoint of the hy_tech resource.
func (c *Client) NewDeleteTechHyTechRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetTechHyTechPath computes a request path to the GetTech action of hy_tech.
func GetTechHyTechPath(techID int) string {
	param0 := strconv.Itoa(techID)

	return fmt.Sprintf("/api/tech/%s", param0)
}

// Retrieve tech with given tech id
func (c *Client) GetTechHyTech(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetTechHyTechRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetTechHyTechRequest create the request corresponding to the GetTech action endpoint of the hy_tech resource.
func (c *Client) NewGetTechHyTechRequest(ctx context.Context, path string) (*http.Request, error) {
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

// TechListHyTechPath computes a request path to the TechList action of hy_tech.
func TechListHyTechPath() string {

	return fmt.Sprintf("/api/tech")
}

// List all techs
func (c *Client) TechListHyTech(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewTechListHyTechRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTechListHyTechRequest create the request corresponding to the TechList action endpoint of the hy_tech resource.
func (c *Client) NewTechListHyTechRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateTechHyTechPayload is the hy_tech UpdateTech action payload.
type UpdateTechHyTechPayload struct {
	// Tech name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// UpdateTechHyTechPath computes a request path to the UpdateTech action of hy_tech.
func UpdateTechHyTechPath(techID int) string {
	param0 := strconv.Itoa(techID)

	return fmt.Sprintf("/api/tech/%s", param0)
}

// Change tech properties
func (c *Client) UpdateTechHyTech(ctx context.Context, path string, payload *UpdateTechHyTechPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateTechHyTechRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTechHyTechRequest create the request corresponding to the UpdateTech action endpoint of the hy_tech resource.
func (c *Client) NewUpdateTechHyTechRequest(ctx context.Context, path string, payload *UpdateTechHyTechPayload, contentType string) (*http.Request, error) {
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
