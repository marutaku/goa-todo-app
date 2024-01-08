package design

import . "goa.design/goa/v3/dsl"

var _ = Service("task", func() {
	Description("The task service manages task lists")
	Error("token_verification_failed", AuthFailedErrorResponse)
	Security(JWT, func() {
		Scope("api:read")
	})
	Method("list", func() {
		Description("List all tasks")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("limit", UInt32, "Maximum number of tasks to return", func() {
				Default(20)
				Example(100)
			})
			Attribute("offset", UInt32, "Offset into the list of tasks to start at", func() {
				Default(0)
				Example(0)
			})
			Attribute("name", String, "Filter by name", func() {
				Default("")
				Example("task1")
			})
		})
		Result(func() {
			Attribute("tasks", CollectionOf(StoredTask), "List of tasks")
		})
		HTTP(func() {
			GET("/tasks")
			Param("limit")
			Param("offset")
			Param("name")
			Response(StatusOK)
			Response("token_verification_failed", StatusUnauthorized)
		})
	})
	Method("show", func() {
		Description("Show a task")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt32, "ID of task to show")
			Required("id")
		})
		Result(func() {
			Attribute("task", StoredTask, "task to show")
		})
		Error("no_match", String, "No task matched given criteria")

		HTTP(func() {
			GET("/tasks/{id}")
			Response(StatusOK)
			Response("no_match", StatusNotFound)
			Response("token_verification_failed", StatusUnauthorized)
		})
	})
	Method("create", func() {
		Description("Create a task")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt32, "ID of task to create")
			Attribute("name", String, "Name of the task")
			Attribute("description", String, "Description of the task")
			Attribute("created_by", String, "Who created the task")
			Required("name", "created_by")
		})
		Result(func() {
			Attribute("task", StoredTask, "Created task")
		})
		HTTP(func() {
			POST("/tasks")
			Response(StatusCreated)
			Response("token_verification_failed", StatusUnauthorized)
		})
	})
	Method("update", func() {
		Description("Update a task")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt32, "ID of task to update")
			Attribute("name", String, "Name of the task")
			Attribute("description", String, "Description of the task")
			Required("id")
		})
		Error("no_match", String, "No task matched given criteria")
		Result(func() {
			Attribute("task", StoredTask, "Updated task")
		})
		HTTP(func() {
			PUT("/tasks/{id}")
			Response(StatusOK)
			Response("token_verification_failed", StatusUnauthorized)
		})
	})
	Method("done", func() {
		Description("Mark a task as done")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt32, "ID of task to mark as done")
			Attribute("done_by", String, "Who did the task")
			Required("id", "done_by")
		})
		Error("no_match", String, "No task matched given criteria")
		Result(func() {
			Attribute("task", StoredTask, "Finished task")
		})
		HTTP(func() {
			PUT("/tasks/{id}/done")
			Response(StatusOK)
			Response("token_verification_failed", StatusUnauthorized)
		})
	})
	Method("delete", func() {
		Description("Delete a task")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("id", UInt32, "ID of task to delete")
			Required("id")
		})
		Error("no_match", String, "No task matched given criteria")
		HTTP(func() {
			DELETE("/tasks/{id}")
			Response(StatusOK)
			Response("token_verification_failed", StatusUnauthorized)
		})
	})
})
