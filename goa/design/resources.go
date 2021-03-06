package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var resourcePrefix string = "hy_"

//Resources group related API endpoints
// This is request object
// Payload is sending data to server

//-----------------------------------------------------------------------------
// User
//-----------------------------------------------------------------------------
var _ = Resource(resourcePrefix+"user", func() {

	DefaultMedia(User) //Response Media Type
	BasePath("/user")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	//Actions define a single API endpoint
	// This name should be unique
	Action("UserList", func() {
		Routing(
			GET(""),
		)
		//NoSecurity()

		Description("Retrieve all users.")

		// No Payload

		//Responses define the shape and status code
		//Response(OK, User) // = Response(OK)
		Response(OK, CollectionOf(User)) //multiple response
		Response(NoContent)
		Response(BadRequest, ErrorMedia)

		//Response(OK, func() {
		//	Media(CollectionOf(User, func() {
		//		View("default")
		//		View("tiny")
		//	}))
		//
		//	// Headers list the response HTTP headers
		//	Headers(func() {
		//		Header("X-Request-Id")
		//	  Required("X-Request-Id")
		//	})
		//})

		//HTTP Request Header
		//Headers(func() {                  // Headers describe relevant action headers
		//    Header("Authorization", String)
		//    Header("X-Account", Integer)
		//    Required("Authorization", "X-Account")
		//})
	})

	Action("GetUser", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id.")

		// Params is for get parameter
		Params(func() {
			Param("userID", Integer, "User ID", func() {
				Minimum(1)
			})
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("CreateUser", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")

		//Payload is for POST data
		//Payload(func() {
		//	Member("name")
		//	Required("name")
		//})

		Payload(UserPayload, func() {
			Required("user_name", "email", "password")
		})

		//Response(Created, "/user/[0-9]+") //[??]
		Response(OK)
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateUser", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Change user properties")

		Params(func() {
			Param("userID", Integer, "User ID")
		})
		//Payload(UserPayload)
		Payload(UserPayload, func() {
			Required("user_name", "email", "password")
		})

		Response(OK)
		//Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("DeleteUser", func() {
		Routing(
			DELETE("/:userID"),
		)
		Description("Delete user ")

		Params(func() {
			Param("userID", Integer, "User ID")
		})

		Response(OK)
		//Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

//-----------------------------------------------------------------------------
// User Tech
//-----------------------------------------------------------------------------
var _ = Resource(resourcePrefix+"usertech", func() {

	DefaultMedia(UserTech) //Response Media Type
	BasePath("/user")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	//
	Action("GetUserLikeTech", func() {
		Routing(
			GET("/:userID/liketech"),
		)
		Description("Retrieve user's favorite techs.")

		// Params is for get parameter
		Params(func() {
			Param("userID", Integer, "User ID", func() {
				Minimum(1)
			})
		})
		Response(OK, CollectionOf(UserTech)) //multiple response
		Response(Unauthorized)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	//
	Action("GetUserDislikeTech", func() {
		Routing(
			GET("/:userID/disliketech"),
		)
		Description("Retrieve user's dislike techs.")

		// Params is for get parameter
		Params(func() {
			Param("userID", Integer, "User ID", func() {
				Minimum(1)
			})
		})

		Response(OK, CollectionOf(UserTech)) //multiple response
		Response(Unauthorized)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

//-----------------------------------------------------------------------------
// User Work History
//-----------------------------------------------------------------------------
var _ = Resource(resourcePrefix+"userWorkHistory", func() {

	DefaultMedia(UserWorkHistory) //Response Media Type
	BasePath("/user")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	//
	Action("GetUserWorkHistory", func() {
		Routing(
			GET("/:userID/workhistory"),
		)
		Description("Retrieve user's work history.")

		// Params is for get parameter
		Params(func() {
			Param("userID", Integer, "User ID", func() {
				Minimum(1)
			})
		})
		Response(OK, CollectionOf(UserWorkHistory)) //multiple response
		Response(Unauthorized)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

//-----------------------------------------------------------------------------
// Tech
//-----------------------------------------------------------------------------
var _ = Resource(resourcePrefix+"tech", func() {

	DefaultMedia(Tech)
	BasePath("/tech")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	// Parent sets the resource parent
	//Parent("tech")  //TODO: how does it work??

	Action("TechList", func() {
		Routing(
			GET(""),
		)
		Description("List all techs")
		//NoSecurity()

		Response(OK, CollectionOf(Tech)) //multiple response
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
	})

	Action("GetTech", func() {
		Routing(
			GET("/:techID"),
		)
		Description("Retrieve tech with given tech id")
		Params(func() {
			Param("techID", Integer, "Tech ID")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("CreateTech", func() {
		Routing(
			POST(""),
		)
		Description("Create new tech")
		Payload(TechPayload, func() {
			//TODO:required value in media_type is given priority over this part...
			Required("name")
		})

		//no response template named "Created" in resource "hy_tech" action "CreateTech"
		//=>it should be defined in api_definition.go
		//Response(Created, "^/tech/[0-9]+$")
		Response(OK)
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateTech", func() {
		Routing(
			PUT("/:techID"),
		)
		Description("Change tech properties")

		Params(func() {
			Param("techID", Integer, "Tech ID")
		})
		Payload(TechPayload, func() {
			Required("name")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("DeleteTech", func() {
		Routing(
			DELETE("/:techID"),
		)
		Description("Delete tech")

		Params(func() {
			Param("techID", Integer, "Tech ID")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

//-----------------------------------------------------------------------------
// Company
//-----------------------------------------------------------------------------
var _ = Resource(resourcePrefix+"company", func() {

	DefaultMedia(Company)
	BasePath("/company")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	// Parent sets the resource parent
	//Parent("user")  //TODO: how does it work??

	Action("CompanyList", func() {
		Routing(
			GET(""),
		)
		Description("List all companies")
		//NoSecurity()

		Response(OK, CollectionOf(Company)) //multiple response
		Response(NoContent)
		Response(BadRequest, ErrorMedia)
	})

	Action("GetCompanyGroup", func() {
		Routing(
			GET("/:companyID"),
		)
		Description("Retrieve company with given company_id")
		Params(func() {
			Param("companyID", Integer, "Company ID")
			//query string
			//Param("hq_flg", String)
			Param("hq_flg", func() {
				Enum("1", "0")
			})
		})
		Response(OK, CollectionOf(Company))
		//Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	//TODO: is it for websocket??
	//Action("watch", func() {
	//	Routing(
	//		GET("/:bottleID/watch"),
	//	)
	//	Scheme("ws")
	//	Description("Retrieve bottle with given id")
	//	Params(func() {
	//		Param("bottleID", Integer)
	//	})
	//	Response(SwitchingProtocols)
	//	Response(BadRequest, ErrorMedia)
	//})

	Action("CreateCompany", func() {
		Routing(
			POST(""),
		)
		Description("Create new company")
		Payload(CompanyPayload, func() {
			//TODO:required value in media_type is given priority over this part...
			Required("name", "country_id", "address")
		})

		//no response template named "Created" in resource "hy_company" action "CreateCompany"
		//=>it should be defined in api_definition.go
		//Response(Created, "^/user/[0-9]+/company/[0-9]+$")
		Response(OK)
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateCompany", func() {
		Routing(
			PUT("/:companyID"),
		)
		Description("Change company properties")

		Params(func() {
			Param("companyID", Integer, "Company ID")
		})
		Payload(CompanyPayload, func() {
			Required("name", "country_id", "address")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("DeleteCompany", func() {
		Routing(
			DELETE("/:companyID"),
		)
		Description("Delete company")

		Params(func() {
			Param("companyID", Integer, "Company ID")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

//-----------------------------------------------------------------------------
// Company branch
//-----------------------------------------------------------------------------
var _ = Resource(resourcePrefix+"companybranch", func() {

	DefaultMedia(Company)
	BasePath("/company/branch")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	// Parent sets the resource parent
	//Parent("user")  //TODO: how does it work??

	Action("GetCompanyBranch", func() {
		Routing(
			GET("/:ID"),
		)
		Description("Retrieve company branch with given id")
		Params(func() {
			//Param("companyID", Integer, "Company ID")
			Param("ID", Integer, "Company detail ID")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("CreateCompanyBranch", func() {
		//TODO:somehow POST path should be unique
		Routing(
			//POST("/:companyID"),
			POST("/:ID"),
		)
		Description("Create new company branch")
		Params(func() {
			//Param("companyID", Integer, "Company ID")
			Param("ID", Integer, "Company ID") //TODO:Though this name is ID, but companyID is set.
		})
		Payload(CompanyTinyPayload, func() {
			//TODO:required value in media_type is given priority over this part...
			Required("country_id", "address")
		})

		//no response template named "Created" in resource "hy_company" action "CreateCompany"
		//=>it should be defined in api_definition.go
		//Response(Created, "^/user/[0-9]+/company/[0-9]+$")
		Response(OK)
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateCompanyBranch", func() {
		Routing(
			PUT("/:ID"),
		)
		Description("Change company branch properties")

		Params(func() {
			//Param("companyID", Integer, "Company ID")
			Param("ID", Integer, "Company detail ID")
		})
		Payload(CompanyTinyPayload, func() {
			Required("country_id", "address")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("DeleteCompanyBranch", func() {
		Routing(
			DELETE("/:ID"),
		)
		Description("Delete company branch")

		Params(func() {
			//Param("companyID", Integer, "Company ID")
			Param("ID", Integer, "Company detail ID")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

//-----------------------------------------------------------------------------
// Public
//-----------------------------------------------------------------------------
var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/*filepath", "public/")
	Files("/swagger-ui/*filepath", "resources/swagger-ui/dist/")
	Files("/swagger.json", "goa/swagger/swagger.json")
})

//var _ = Resource("public", func() {
//	Origin("*", func() {
//		Methods("GET", "OPTIONS")
//	})
//	Files("/", "public/index.html")
//})

//var _ = Resource("js", func() {
//	Origin("*", func() {
//		Methods("GET", "OPTIONS")
//	})
//	Files("/js/*filepath", "public/js")
//})

//var _ = Resource("swagger", func() {
//	Origin("*", func() {
//		Methods("GET", "OPTIONS")
//	})
//	Files("/swagger.json", "public/swagger/swagger.json")
//})
