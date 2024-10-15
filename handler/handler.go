package handler

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/reiver/go-telnet"
)

func CheckPort(ctx *fiber.Ctx) error {

	var ip = ctx.Params("ip")
	var port = ctx.Params("port")
	adress := net.JoinHostPort(ip,port)

	conn,err := telnet.DialTo(adress)

	if err != nil {
		fmt.Println("error fail false")
		return ctx.SendString("Port is passive")

	
	}

	conn.Close()
	return ctx.SendString("Port is active")

}