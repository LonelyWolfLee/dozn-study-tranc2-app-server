package account

import "github.com/gofiber/fiber/v2"

func create(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : create")
}

func list(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : list")
}
