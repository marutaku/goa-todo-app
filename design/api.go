package design

import (
	. "goa.design/goa/v3/dsl"
)

var JWT = JWTSecurity("jwt", func() {
	Description("JWT Authentication for the API")
	Scope("api:read", "Provides read access")
	Scope("api:write", "Provides write access")
})

var Task = Type("Task", func() {
	Description("Task describes a todo item")

	Attribute("id", UInt32, "Unique ID", func() {
		Meta("struct:tag:gorm", "primaryKey")
		Meta("struct:tag:json", "id")
	})
	Attribute("name", String, "Name of the todo", func() {
		Meta("struct:tag:gorm", "not null")
		Meta("struct:tag:json", "name")
	})
	Attribute("description", String, "Description of the todo", func() {
		Meta("struct:tag:gorm", "not null;default ''")
		Meta("struct:tag:json", "description")
	})
	Attribute("done", Boolean, "Whether or not the todo is done", func() {
		Default(false)
		Meta("struct:tag:gorm", "not null;default false")
		Meta("struct:tag:json", "done")
	})
	Attribute("doneAt", String, "When the todo was done in ISO format", func() {
		Default("")
		Meta("struct:tag:gorm", "default ''")
		Meta("struct:tag:json", "doneAt")
	})
	Attribute("doneBy", String, "Who did the todo", func() {
		Default("")
		Meta("struct:tag:gorm", "default ''")
		Meta("struct:tag:json", "doneBy")
	})
	Attribute("createdAt", String, "When the todo was created in ISO format", func() {
		Meta("struct:tag:gorm", "not null")
		Meta("struct:tag:json", "createdAt")
	})
	Attribute("createdBy", String, "Who created the todo", func() {
		Meta("struct:tag:gorm", "not null")
		Meta("struct:tag:json", "createdBy")
	})

	Required("id", "name", "description", "done", "createdAt", "createdBy")
})

var StoredTask = ResultType("applicaiton/vnd.backend.stored-task", func() {
	Description("A task")
	Reference(Task)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("done")
		Attribute("doneAt")
		Attribute("doneBy")
		Attribute("createdAt", func() {

		})
		Attribute("createdBy")
	})
})

var _ = API("Task", func() {
	Title("Task Service")
	Error("token_verification_failed", AuthFailedErrorResponse, "Token verification failed")
	Description("Service for managing tasks")
	Server("task", func() {
		Host("local", func() {
			URI("http://0.0.0.0:8000")
		})
	})
})
