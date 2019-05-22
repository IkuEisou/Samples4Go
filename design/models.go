package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("CongoStorageGroup", func() {
	Description("This is the global storage group")
	Store("MySQL", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("User", func() {
			// BuildsFrom(func() {
			// Payload("user", "create")
			// Payload("user", "update")
			// })
			RendersTo(User)
			Description("User Model Description")
			HasMany("Books", "Book")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Book", func() {
			// BuildsFrom(func() {
			// Payload("book", "create")
			// Payload("book", "update")
			// })
			RendersTo(Book)
			Description("Book Model")
			BelongsTo("User")
			// HasMany("Reviews", "Review")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Payload Model PK field")
			})
			Field("title", func() {
				Alias("proposal_title")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Test", func() {
			Description("TestModel")
			NoAutomaticIDFields()
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
