package main

import (
	"github.com/ruchitdobariya/gosecAPI/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{})

	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	routers.SetRouters(app)

	app.Listen(":3000")
}
