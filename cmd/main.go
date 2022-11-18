package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/melinaco4/c-handler/pkg/controllers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", controllers.status)

	app.Get("/companies", controllers.GetCompanies)
	app.Post("/company", controllers.CreateCompany)
	app.Patch("/company/:id", controllers.UpdateCompany)
	app.Delete("/company/:id", controllers.DeleteCompany)
}

func main() {
	if err := controllers.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
