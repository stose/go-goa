// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewXMLDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Login(*LoginAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/auth/login", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*LoginAuthPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/api/auth/login", ctrl.MuxHandler("Login", h, unmarshalLoginAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Login", "route", "POST /api/auth/login")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalLoginAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &loginAuthPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Health(*HealthHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/_ah/health", ctrl.MuxHandler("preflight", handleHealthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	h = handleHealthOrigin(h)
	service.Mux.Handle("GET", "/api/_ah/health", ctrl.MuxHandler("health", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Health", "route", "GET /api/_ah/health")
}

// handleHealthOrigin applies the CORS response headers corresponding to the origin.
func handleHealthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// HyCompanyController is the controller interface for the HyCompany actions.
type HyCompanyController interface {
	goa.Muxer
	CompanyList(*CompanyListHyCompanyContext) error
	CreateCompany(*CreateCompanyHyCompanyContext) error
	CreateCompanyBranch(*CreateCompanyBranchHyCompanyContext) error
	DeleteCompany(*DeleteCompanyHyCompanyContext) error
	GetCompanyBranch(*GetCompanyBranchHyCompanyContext) error
	GetCompanyGroup(*GetCompanyGroupHyCompanyContext) error
	UpdateCompany(*UpdateCompanyHyCompanyContext) error
}

// MountHyCompanyController "mounts" a HyCompany resource controller on the given service.
func MountHyCompanyController(service *goa.Service, ctrl HyCompanyController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/company", ctrl.MuxHandler("preflight", handleHyCompanyOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/company/:companyID/branch/", ctrl.MuxHandler("preflight", handleHyCompanyOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/company/:companyID", ctrl.MuxHandler("preflight", handleHyCompanyOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/company/:companyID/branch/:ID", ctrl.MuxHandler("preflight", handleHyCompanyOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCompanyListHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.CompanyList(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("GET", "/api/company", ctrl.MuxHandler("CompanyList", h, nil))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "CompanyList", "route", "GET /api/company", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCompanyHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCompanyHyCompanyPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateCompany(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("POST", "/api/company", ctrl.MuxHandler("CreateCompany", h, unmarshalCreateCompanyHyCompanyPayload))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "CreateCompany", "route", "POST /api/company", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCompanyBranchHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCompanyBranchHyCompanyPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateCompanyBranch(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("POST", "/api/company/:companyID/branch/", ctrl.MuxHandler("CreateCompanyBranch", h, unmarshalCreateCompanyBranchHyCompanyPayload))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "CreateCompanyBranch", "route", "POST /api/company/:companyID/branch/", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCompanyHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteCompany(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("DELETE", "/api/company/:companyID", ctrl.MuxHandler("DeleteCompany", h, nil))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "DeleteCompany", "route", "DELETE /api/company/:companyID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetCompanyBranchHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetCompanyBranch(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("GET", "/api/company/:companyID/branch/:ID", ctrl.MuxHandler("GetCompanyBranch", h, nil))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "GetCompanyBranch", "route", "GET /api/company/:companyID/branch/:ID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetCompanyGroupHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetCompanyGroup(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("GET", "/api/company/:companyID", ctrl.MuxHandler("GetCompanyGroup", h, nil))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "GetCompanyGroup", "route", "GET /api/company/:companyID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCompanyHyCompanyContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateCompanyHyCompanyPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateCompany(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyCompanyOrigin(h)
	service.Mux.Handle("PUT", "/api/company/:companyID", ctrl.MuxHandler("UpdateCompany", h, unmarshalUpdateCompanyHyCompanyPayload))
	service.LogInfo("mount", "ctrl", "HyCompany", "action", "UpdateCompany", "route", "PUT /api/company/:companyID", "security", "jwt")
}

// handleHyCompanyOrigin applies the CORS response headers corresponding to the origin.
func handleHyCompanyOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateCompanyHyCompanyPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCompanyHyCompanyPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCompanyHyCompanyPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreateCompanyBranchHyCompanyPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCompanyBranchHyCompanyPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCompanyBranchHyCompanyPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateCompanyHyCompanyPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateCompanyHyCompanyPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateCompanyHyCompanyPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HyUserController is the controller interface for the HyUser actions.
type HyUserController interface {
	goa.Muxer
	CreateUser(*CreateUserHyUserContext) error
	DeleteUser(*DeleteUserHyUserContext) error
	GetUser(*GetUserHyUserContext) error
	UpdateUser(*UpdateUserHyUserContext) error
	UserList(*UserListHyUserContext) error
}

// MountHyUserController "mounts" a HyUser resource controller on the given service.
func MountHyUserController(service *goa.Service, ctrl HyUserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/user", ctrl.MuxHandler("preflight", handleHyUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/user/:userID", ctrl.MuxHandler("preflight", handleHyUserOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserHyUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserHyUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyUserOrigin(h)
	service.Mux.Handle("POST", "/api/user", ctrl.MuxHandler("CreateUser", h, unmarshalCreateUserHyUserPayload))
	service.LogInfo("mount", "ctrl", "HyUser", "action", "CreateUser", "route", "POST /api/user", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteUserHyUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyUserOrigin(h)
	service.Mux.Handle("DELETE", "/api/user/:userID", ctrl.MuxHandler("DeleteUser", h, nil))
	service.LogInfo("mount", "ctrl", "HyUser", "action", "DeleteUser", "route", "DELETE /api/user/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserHyUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyUserOrigin(h)
	service.Mux.Handle("GET", "/api/user/:userID", ctrl.MuxHandler("GetUser", h, nil))
	service.LogInfo("mount", "ctrl", "HyUser", "action", "GetUser", "route", "GET /api/user/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUserHyUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserHyUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateUser(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyUserOrigin(h)
	service.Mux.Handle("PUT", "/api/user/:userID", ctrl.MuxHandler("UpdateUser", h, unmarshalUpdateUserHyUserPayload))
	service.LogInfo("mount", "ctrl", "HyUser", "action", "UpdateUser", "route", "PUT /api/user/:userID", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUserListHyUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.UserList(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleHyUserOrigin(h)
	service.Mux.Handle("GET", "/api/user", ctrl.MuxHandler("UserList", h, nil))
	service.LogInfo("mount", "ctrl", "HyUser", "action", "UserList", "route", "GET /api/user", "security", "jwt")
}

// handleHyUserOrigin applies the CORS response headers corresponding to the origin.
func handleHyUserOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateUserHyUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserHyUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createUserHyUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserHyUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserHyUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUserHyUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/*filepath", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger-ui/*filepath", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "goa/swagger/swagger.json")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "goa/swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/*filepath", "public/")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/", "route", "GET /*filepath")

	h = ctrl.FileHandler("/swagger-ui/*filepath", "resources/swagger-ui/dist/")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "resources/swagger-ui/dist/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/", "public/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/index.html", "route", "GET /")

	h = ctrl.FileHandler("/swagger-ui/", "resources/swagger-ui/dist/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "resources/swagger-ui/dist/index.html", "route", "GET /swagger-ui/")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://swagger.goa.design") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
