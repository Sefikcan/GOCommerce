package routes

import (
	. "basket/api"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
	app.Post("/api/v1/baskets",
		AddOrUpdateBasket)

	app.Get("/api/v1/baskets/:userId",
		GetBasketByUserId)

	app.Delete("/api/v1/baskets/:userId",
		RemoveBasketByUserId)
}
