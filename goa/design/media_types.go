package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Define Response data
// It contains the media type data structures used by resource actions to build the responses
// Add variable and set MediaType object

// reference
// https://goa.design/reference/goa/design/apidsl/#func-mediatype-a-name-apidsl-mediatype-a

//-----------------------------------------------------------------------------
// Authorized is the authority resource media type.
//-----------------------------------------------------------------------------
var Authorized = MediaType("application/vnd.authorized+json", func() {
	// Response Description
	Description("An authorized response")

	Attributes(func() {
		Attribute("token", String, "JWT token", func() {
			Example("token.string")
		})
		Attribute("id", Integer, "User ID", fieldID)
		Required("token", "id")
	})

	//default MediaType
	View("default", func() {
		Attribute("token")
		Attribute("id")
	})
})

//-----------------------------------------------------------------------------
// User is the user resource media type.
//-----------------------------------------------------------------------------
var User = MediaType("application/vnd.user+json", func() {
	// Response Description
	Description("A user information")

	ContentType("application/json")

	// Reference can be used in: MediaType, Type
	// Though set Reference, Attribute is required
	Reference(UserPayload)
	Reference(CommonResponse)

	//`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
	//`user_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'User Name',
	//`email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
	//`password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
	//`delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
	//`created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
	//`updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',

	Attributes(func() {
		Attribute("id", Integer, "User ID", fieldID)
		Attribute("user_name")
		Attribute("email")
		Attribute("password")
		Attribute("created_at")
		Attribute("updated_at")

		//TODO:required value in media_type is given priority over resource...
		//That means, this part would affect on Action-Payload-Required
		//when field is zero or empty, data is not return unless it's not set in Required
		Required("user_name", "email", "password")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	//View default is for UserList, GetUser,
	View("default", func() {
		Attribute("id")
		Attribute("user_name")
		Attribute("email")
	})

	View("id", func() {
		Description("id is the view used for C U D")
		Attribute("id")
	})
})

//-----------------------------------------------------------------------------
// UserTech is the user tech resource media type.
//-----------------------------------------------------------------------------
var UserTech = MediaType("application/vnd.usertech+json", func() {
	// Response Description
	Description("A user information")

	ContentType("application/json")

	// Reference can be used in: MediaType, Type
	// Though set Reference, Attribute is required
	Reference(UserTechPayload)

	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
	//`user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
	//`tech_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech ID',
	//`name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech Name',

	Attributes(func() {
		Attribute("id", Integer, "Tech ID", fieldID)
		Attribute("tech_name")

		//TODO:required value in media_type is given priority over resource...
		//That means, this part would affect on Action-Payload-Required
		//when field is zero or empty, data is not return unless it's not set in Required
		Required("tech_name")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	//View default is for UserList, GetUser,
	View("default", func() {
		Attribute("id")
		Attribute("tech_name")
	})

	View("tech", func() {
		Description("id is the view used for C U D")
		Attribute("tech_name")
	})
})

//-----------------------------------------------------------------------------
// UserTech is the user tech resource media type.
//-----------------------------------------------------------------------------
var UserWorkHistory = MediaType("application/vnd.userworkhistory+json", func() {
	// Response Description
	Description("A user information")

	ContentType("application/json")

	// Reference can be used in: MediaType, Type
	// Though set Reference, Attribute is required
	Reference(UserWorkHistoryPayload)

	//`id`          int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
	//`user_id`     int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
	//`company_branch_id`  int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company Branch ID',
	//`title`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Title',
	//`description` json NOT NULL COMMENT'Description',
	//`started_at`  date DEFAULT NULL COMMENT'Started Date',
	//`ended_at`    date DEFAULT NULL COMMENT'Ended Date',
	//`delete_flg`  char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
	//`created_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
	//`updated_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',

	Attributes(func() {
		Attribute("title")
		Attribute("company")
		Attribute("country")
		Attribute("term")
		Attribute("description")
		Attribute("techs")

		//TODO:required value in media_type is given priority over resource...
		//That means, this part would affect on Action-Payload-Required
		//when field is zero or empty, data is not return unless it's not set in Required
		Required("title", "company", "country")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	//View default is for UserList, GetUser,
	View("default", func() {
		Attribute("title")
		Attribute("company")
		Attribute("country")
		Attribute("term")
		Attribute("description")
		Attribute("techs")
	})
})

//-----------------------------------------------------------------------------
// Tech is the tech resource media type.
//-----------------------------------------------------------------------------
var Tech = MediaType("application/vnd.tech+json", func() {
	Description("A tech information")

	ContentType("application/json")

	Reference(TechPayload)
	Reference(CommonResponse)

	Attributes(func() {
		Attribute("id", Integer, "Tech ID", fieldID)
		Attribute("name")
		Attribute("created_at")
		Attribute("updated_at")

		//TODO:required value in media_type is given priority over resource...
		//That means, this part would affect on Action-Payload-Required
		//when field is zero or empty, data is not return unless it's not set in Required
		Required("name")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})

	View("id", func() {
		Description("id is the view used for C U D")
		Attribute("id")
	})
})

//-----------------------------------------------------------------------------
// Company is the company resource media type.
//-----------------------------------------------------------------------------
var Company = MediaType("application/vnd.company+json", func() {
	Description("A company information")

	ContentType("application/json")

	Reference(CompanyPayload)
	Reference(CommonResponse)

	Attributes(func() {
		Attribute("id", Integer, "Company Detail ID", fieldID)
		Attribute("company_id", Integer, "Company ID", fieldID)
		Attribute("name")
		Attribute("hq_flg")
		Attribute("country_name")
		Attribute("address")
		Attribute("created_at")
		Attribute("updated_at")

		//TODO:required value in media_type is given priority over resource...
		//That means, this part would affect on Action-Payload-Required
		//when field is zero or empty, data is not return unless it's not set in Required
		Required("name", "address")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	View("default", func() {
		Attribute("id")
		Attribute("company_id")
		Attribute("name")
		Attribute("hq_flg")
		Attribute("country_name")
		Attribute("address")
		//Attribute("created_at")
		//Attribute("created_by")
	})
	View("detailid", func() {
		Description("only company's detail id")
		Attribute("id")
	})

	View("id", func() {
		Description("only company's id")
		Attribute("company_id")
	})

	View("idname", func() {
		Description("only company's id and name")
		Attribute("company_id")
		Attribute("name")
	})
})

//-----------------------------------------------------------------------------
// UserCompany is the company resource media type.
//-----------------------------------------------------------------------------
//TODO: unimplemented
var UserCompany = MediaType("application/vnd.usercomany+json", func() {
	Description("A user who belongs to which companies")

	// So define types and take advantage of the Reference keyword
	// when the same attributes are shared between types and media types
	//Reference(BottlePayload)

	Attributes(func() {
		Attribute("id", Integer, "ID of user company", func() {
			Example(1)
		})
		Attribute("href", String, "API href of bottle", func() {
			Example("/user/1/company")
		})

		Attribute("user_id", Integer, "ID of user id", func() {
			Example(1)
		})

		Attribute("company_id", Integer, "ID of user id", func() {
			Example(1)
		})

		//Attribute("rating", Integer, "Rating of bottle between 1 and 5", func() {
		//	Minimum(1)
		//	Maximum(5)
		//})

		//From User object
		//Attribute("user", User, "User who belong to company")
		//From Company object
		//Attribute("company", Company, "user who work(ed) for company")

		Reference(CommonResponse)

		Required("id", "href", "user_id", "company_id")
		//Required("created_at")
	})

	//Links are rendered using the special "link" view.
	// Media types that are linked to must define that view. Here is an example
	// showing all the possible media type sub-definitions:

	//When including other media type from outside, this link would be required.

	//Links(func() {
	//	Link("user")
	//})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("user_id")
		Attribute("company_id")

		//Attribute("account", func() {
		//	View("tiny")
		//})
		//Attribute("links")
	})
})
