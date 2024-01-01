package design

import (
	. "goa.design/goa/v3/dsl"
)

var Task = Type("Task", func() {
	Description("Task describes a todo item")

	Attribute("id", UInt32, "Unique ID", func() {
		Meta("struct:tag:gorm", "primaryKey")
	})
	Attribute("name", String, "Name of the todo", func() {
		Meta("struct:tag:gorm", "not null")
	})
	Attribute("description", String, "Description of the todo", func() {
		Meta("struct:tag:gorm", "not null;default ''")
	})
	Attribute("done", Boolean, "Whether or not the todo is done", func() {
		Default(false)
		Meta("struct:tag:gorm", "not null;default false")
	})
	Attribute("doneAt", String, "When the todo was done in ISO format", func() {
		Default("")
		Meta("struct:tag:gorm", "default ''")
	})
	Attribute("doneBy", String, "Who did the todo", func() {
		Default("")
		Meta("struct:tag:gorm", "default ''")
	})
	Attribute("createdAt", String, "When the todo was created in ISO format", func() {
		Meta("struct:tag:gorm", "not null")
	})
	Attribute("createdBy", String, "Who created the todo", func() {
		Meta("struct:tag:gorm", "not null")
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
	Description("Service for managing tasks")
	Server("task", func() {
		Host("local", func() {
			URI("http://0.0.0.0:8000")
		})
	})
})
