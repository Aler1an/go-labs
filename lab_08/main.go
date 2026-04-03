package main

import "github.com/gofiber/fiber/v3"

func main() {
	connectDB()

	app := fiber.New()

	app.Get("/contacts", getAllContacts)
	app.Get("/contacts/:id", getContactByID)
	app.Post("/contacts", createContact)
	app.Put("/contacts/:id", updateContact)
	app.Delete("/contacts/:id", deleteContact)

	app.Listen(":3000")
}
