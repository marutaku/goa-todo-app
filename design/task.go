package design

import . "goa.design/goa/v3/dsl"

var _ = Service("task", func() {
	Description("The task service manages task lists")
	Method("list", func() {
		Description("List all tasks")
		Payload(func() {
			Attribute("limit", UInt32, "Maximum number of tasks to return", func() {
				Default(20)
			})
			Attribute("offset", UInt32, "Offset into the list of tasks to start at", func() {
				Default(0)
			})
		})
		Result(func() {
			Attribute("tasks", CollectionOf(StoredTask), "List of tasks")
		})
		HTTP(func() {
			GET("/tasks")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show a task")
		Payload(func() {
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
		})
	})
})
