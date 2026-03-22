package main

import (
	"strconv"

	"github.com/Aler1an/go-labs/lab_07/models"

	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
)

func getAllNotes(c fiber.Ctx) error {
	raw, err := easyjson.Marshal(notes)
	if err != nil {
		return c.Status(500).SendString("Error")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(raw)
}

func getNoteByID(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for _, note := range notes {
		if note.ID == id {
			return c.JSON(note)
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
}

func createNote(c fiber.Ctx) error {
	var note models.Note

	if err := easyjson.Unmarshal(c.Body(), &note); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if note.ID == 0 {
		note.ID = nextID
		nextID++
	} else {
		for _, n := range notes {
			if n.ID == note.ID {
				return c.Status(409).JSON(fiber.Map{"error": "ID already exists"})
			}
		}
	}

	if note.ID >= nextID {
		nextID = note.ID + 1
	}

	notes = append(notes, note)
	return c.Status(201).JSON(note)
}

func updateNote(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updatedData models.Note

	if err := easyjson.Unmarshal(c.Body(), &updatedData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	for i, note := range notes {
		if note.ID == id {
			updatedData.ID = id
			notes[i] = updatedData
			return c.JSON(updatedData)
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
}

func deleteNote(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
}
