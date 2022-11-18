package routes

import (
	"github.com/gofiber/fiber"
	"github.com/melinaco4/c-handler/pkg/controllers"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/companies", controllers.GetCompanies)
	app.Post("/company", controllers.CreateCompany)
	app.Patch("/company/:id", controllers.UpdateCompany)
	app.Delete("/company/:id", controllers.DeleteCompany)
}
