// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "api": Client
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the api service client.
type Client struct {
	*goaclient.Client
	JWTSigner goaclient.Signer
	Encoder   *goa.HTTPEncoder
	Decoder   *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewXMLDecoder, "*/*")

	return client
}

// SetJWTSigner sets the request signer for the jwt security scheme.
func (c *Client) SetJWTSigner(signer goaclient.Signer) {
	c.JWTSigner = signer
}
