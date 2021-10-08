package routes

import (
	"github.com/gofiber/fiber/v2"
	. "ordering/api"
)

func Setup(app *fiber.App){
	app.Get("/api/v1/orders/:userId",
		GetOrderByUserId)

	app.Post("/api/v1/orders",
		CreateOrder)

	app.Delete("/api/v1/orders/:id",
		DeleteOrder)

	app.Put("/api/v1/orders/:id",
		UpdateOrder)
}
