package main

import (
	"dozn/app-server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	utils.LOG.Trace.Println("Starting the application...")
	utils.LOG.Info.Println("Something noteworthy happened...")
	utils.LOG.Warn.Println("There is something you should know about...")
	utils.LOG.Error.Println("Something went wrong...")

	// Echo instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routes
	// 1. Registration (User)
	// 2. Login (User)
	// 3. Auto login (User)
	// 4. Create bank account (User, Account)
	// 5. Get bank accounts (User, Account)
	// 6. Transaction list for one bank account (User, Account, Transaction)
	// 7. Money tranfer (User, Account, Transaction)
	// 8. Deposit (User, Account, Transaction)
	app.Get("/", hello)

	// Start server
	app.Listen(":3000")
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
