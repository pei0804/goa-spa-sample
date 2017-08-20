package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa-spa", func() {
	Title("Sample")
	Description("Sample")
	Host("localhost:8080")
	Scheme("http", "https")
	BasePath("/")
})

var _ = Resource("home", func() {
	Action("home", func() {
		Description("")
		Routing(GET("/api/home"))
		Params(func() {
			Param("id", Integer, "id")
		})
		Response(OK, HashOf(String, String))
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Files("/swagger/*filepath", "../public/swagger/")
})

var _ = Resource("schema", func() {
	Files("/schema/*filepath", "../public/schema/")
})

var _ = Resource("ui", func() {
	Files("/favicon.ico", "../front/build/favicon.ico")
	Files("/static/*filepath", "../front/build/static")
	Files("/manifest.json", "../front/build/manifest.json")
	Files("*", "../front/build/index.html")
	Files("/", "../front/build/index.html")
})
