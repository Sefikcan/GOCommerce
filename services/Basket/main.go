package main

import (
	"basket/common/swagger"
	Connection "basket/infrastructure"
	"basket/routes"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	Connection.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowOrigins: "*",
	}))

	swaggerMiddleware := swagger.HandleSwagger("./docs/swagger.json", "/")

	swaggerMiddleware.AddSwagger(app)

	prometheus := fiberprometheus.New("gocommerce-basket-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	routes.Setup(app)

	app.Listen(":8002")
}
