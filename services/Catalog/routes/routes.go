package routes

import (
	. "catalog/api"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get(
		"/api/v1/catalogs",
		GetProducts,
	)
	app.Get("/api/v1/catalogs/:id",
		GetProductById)

	app.Post("/api/v1/catalogs",
		CreateProduct)

	app.Delete("/api/v1/catalogs/:id",
		DeleteProduct)

	app.Put("/api/v1/catalogs/:id",
		UpdateProduct)
}
