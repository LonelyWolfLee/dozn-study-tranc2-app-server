package account

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func create(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : create")
}

func list(context *fiber.Ctx) error {
	path := strings.Join(strings.Split(context.Path(), "/")[3:], "/")

	doRequest("http://localhost:3001/"+path, context.Method(), context.Body())

	return context.Next()
}
