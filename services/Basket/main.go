package main

import (
	Connection "basket/infrastructure"
	"basket/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ansrivas/fiberprometheus/v2"
)

func main() {
	Connection.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	prometheus := fiberprometheus.New("gocommerce-basket-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	routes.Setup(app)

	app.Listen(":8002")
}
