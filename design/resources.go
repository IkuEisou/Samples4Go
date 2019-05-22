package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	Action("list", func() {
		Routing(GET("user"))
		Description("Show the users")
		Response(OK, CollectionOf(User))
		// 本番では404返しちゃう
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("book", func() {
	Action("list", func() {
		Routing(GET("book"))
		Description("Show the books")
		Response(OK, CollectionOf(Book))
		// 本番では404返しちゃう
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})
var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "text/plain")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
})
