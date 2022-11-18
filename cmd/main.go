package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/melinaco4/c-handler/pkg/controllers"
)

func main() {
	if err := controllers.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	routes.setupRoutes(app)
}
