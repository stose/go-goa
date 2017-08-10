// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": hy_userWorkHistory Resource Client
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetUserWorkHistoryHyUserWorkHistoryPath computes a request path to the GetUserWorkHistory action of hy_userWorkHistory.
func GetUserWorkHistoryHyUserWorkHistoryPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/api/user/%s/workhistory", param0)
}

// Retrieve user's work history.
func (c *Client) GetUserWorkHistoryHyUserWorkHistory(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserWorkHistoryHyUserWorkHistoryRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserWorkHistoryHyUserWorkHistoryRequest create the request corresponding to the GetUserWorkHistory action endpoint of the hy_userWorkHistory resource.
func (c *Client) NewGetUserWorkHistoryHyUserWorkHistoryRequest(ctx context.Context, path string) (*http.Request, error) {
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