package main

import "github.com/gofiber/fiber/v3"

func main() {

	app := fiber.New()

	app.Get("/notes", getAllNotes)
	app.Get("/notes/:id", getNoteByID)
	app.Post("/notes", createNote)
	app.Put("/notes/:id", updateNote)
	app.Delete("/notes/:id", deleteNote)

	app.Listen(":3000")
}
