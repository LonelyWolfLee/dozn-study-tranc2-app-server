package account

import "github.com/gofiber/fiber/v2"

func create(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : create")
}

func getOne(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : get one")
}

func list(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : list")
}
