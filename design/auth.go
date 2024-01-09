package design

import . "goa.design/goa/v3/dsl"

var AuthFailedErrorResponse = Type("auth_failed", func() {
	Attribute("message", String, "Error message")
	Required("message")
})

var _ = Service("auth", func() {
	Description("The auth service manages authentication")
	Method("login", func() {
		Description("Login to the system")
		Payload(func() {
			Attribute("username", String, "Username to login with")
			Attribute("password", String, "Password to login with")
			Required("username", "password")
		})
		Error("login_failed", String, "User not found")
		Result(func() {
			Attribute("token", String, "JWT token to use for authentication")
			Required("token")
		})
		HTTP(func() {
			POST("/auth/login")
			Response(StatusOK)
			Response("login_failed", StatusUnauthorized)
		})
	})
	Method("register", func() {
		Description("Register a new user")
		Payload(func() {
			Attribute("username", String, "Username to register with")
			Attribute("password", String, "Password to register with")
			Required("username", "password")
		})
		Error("register_failed", String, "Username already exists")
		Result(func() {
			Attribute("token", String, "JWT token to use for authentication")
			Required("token")
		})
		HTTP(func() {
			POST("/auth/register")
			Response(StatusOK)
			Response("register_failed", StatusBadRequest, func() {
				Description("Username already exists")
			})
		})
	})
	// Method("logout", func() {
	// 	Description("Logout of the system")
	// 	Payload(func() {
	// 		Attribute("token", String, "JWT token to use for authentication")
	// 		Required("token")
	// 	})
	// 	HTTP(func() {
	// 		POST("/logout")
	// 		Response(StatusOK)
	// 	})
	// })
})
