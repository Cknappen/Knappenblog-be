package api

import (
	"github.com/Cknappen/Knappenblog-be/src/pkg/api/user"
	"github.com/Cknappen/Knappenblog-be/src/pkg/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ConfigureDatabase(db *gorm.DB) {
	// Migrate the User struct to create the corresponding table in the database
	db.AutoMigrate(&user.User{})
}

func SetupRoutes(app *fiber.App) {
	db := database.InitDB()
	ConfigureDatabase(db)

	app.Use(func(c *fiber.Ctx) error {
		database.SetLocal[*gorm.DB](c, "db", db)
		// Go to next middleware:
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber API!")
	})
	// Serve Swagger UI
	// app.Get("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	todoGroup := app.Group("/todos")
	SetupTodoRoutes(todoGroup)
}
