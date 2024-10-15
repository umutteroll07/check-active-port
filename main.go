package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umutteroll07/telnet_project/route"
)

func main(){

	app := fiber.New()

	route.SetupRoute(app)

	app.Listen(":4000")
}