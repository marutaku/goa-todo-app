package design

import . "goa.design/goa/v3/dsl"

var _ = Service("todo", func() {
	Description("The todo service manages todo lists")
	Method("list", func() {
		Description("List all todos")
		Payload(func() {
			Attribute("limit", UInt32, "Maximum number of todos to return", func() {
				Default(20)
			})
			Attribute("offset", UInt32, "Offset into the list of todos to start at", func() {
				Default(0)
			})
		})
		Result(func() {
			Attribute("todos", CollectionOf(Todo), "List of todos")
		})
		HTTP(func() {
			GET("/todos")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show a todo")
		Payload(func() {
			Attribute("id", UInt32, "ID of todo to show")
			Required("id")
		})
		Result(func() {
			Attribute("todo", Todo, "Todo to show")
		})

		HTTP(func() {
			GET("/todos/{id}")
			Response(StatusOK)
		})
	})
})
