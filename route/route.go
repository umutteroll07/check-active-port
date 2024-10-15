package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umutteroll07/telnet_project/handler"
)

func SetupRoute(app *fiber.App){

	app.Get("/address/:ip/:port",handler.CheckPort)

}