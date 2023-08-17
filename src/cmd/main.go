package main

import (
	"log"

	"github.com/Cknappen/Knappenblog-fe/pkg/api"
	"github.com/Cknappen/Knappenblog-fe/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api.SetupRoutes(app)
	app.Use(auth.AuthenticateJWT)

	log.Fatal(app.Listen(":3000"))
}
