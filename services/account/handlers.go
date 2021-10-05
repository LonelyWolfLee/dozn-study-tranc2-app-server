package account

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func create(context *fiber.Ctx) error {
	return context.SendString("ACCOUNT : create")
}

func list(context *fiber.Ctx) error {
	doRequest("http://localhost:3001"+context.Path(), context.Method(), string(context.Body()[:]))

	var data []byte
	json.Unmarshal(context.Body(), &data)
	fmt.Println(data)

	return context.Next()
}
