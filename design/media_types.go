package design

import (
	"time"

	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

// UserMedia ユーザー詳細情報
var User = MediaType("application/vnd.user+json", func() {
	Description("User")
	Attribute("id", Integer, "user id", func() {
		Example(123)
	})
	Attribute("name", String, "name", func() {
		Example("James Bond")
	})
	Attribute("email", String, "mail address", func() {
		Example("example@gmail.com")
	})
	Attribute("birth", DateTime, "birthday", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(1980, 1, 1, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("books", ArrayOf(Book))
	Required("id", "name", "email", "birth", "books")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("birth")
		Attribute("books")
	})
})

var Book = MediaType("application/vnd.book+json", func() {
	Description("Book")
	Attribute("id", Integer, "book id", func() {
		Example(001)
	})
	Attribute("name", String, "book name", func() {
		Example("Bible")
	})
	Attribute("uid", Integer, "borrower id", func() {
		Example(123)
	})
	Attribute("borrowed_at", DateTime, "borrowed time", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(1980, 1, 1, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("returned_at", DateTime, "returned time", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(1980, 1, 1, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Required("id", "name", "uid", "borrowed_at", "returned_at")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("uid")
		Attribute("borrowed_at")
		Attribute("returned_at")
	})
})
