package main

import (
	"strconv"

	"github.com/Aler1an/go-labs/lab_08/models"

	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
)

func getAllContacts(c fiber.Ctx) error {
	rows, err := DB.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return c.Status(500).SendString("DB error")
	}
	defer rows.Close()

	var contacts models.ContactList

	for rows.Next() {
		var contact models.Contact
		rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		contacts = append(contacts, contact)
	}

	raw, _ := easyjson.Marshal(contacts)

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(raw)
}

func getContactByID(c fiber.Ctx) error {
	id := c.Params("id")

	var contact models.Contact

	err := DB.QueryRow(
		"SELECT id, name, phone FROM contacts WHERE id=$1", id,
	).Scan(&contact.ID, &contact.Name, &contact.Phone)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}

	return c.JSON(contact)
}

func createContact(c fiber.Ctx) error {
	var contact models.Contact

	if err := easyjson.Unmarshal(c.Body(), &contact); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	err := DB.QueryRow(
		"INSERT INTO contacts(name, phone) VALUES($1, $2) RETURNING id",
		contact.Name,
		contact.Phone,
	).Scan(&contact.ID)

	if err != nil {
		return c.Status(500).SendString("Insert error")
	}

	return c.Status(201).JSON(contact)
}

func updateContact(c fiber.Ctx) error {
	id := c.Params("id")

	var contact models.Contact

	if err := easyjson.Unmarshal(c.Body(), &contact); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	res, err := DB.Exec(
		"UPDATE contacts SET name=$1, phone=$2 WHERE id=$3",
		contact.Name,
		contact.Phone,
		id,
	)

	if err != nil {
		return c.Status(500).SendString("Update error")
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}

	contact.ID, _ = strconv.Atoi(id)
	return c.JSON(contact)
}

func deleteContact(c fiber.Ctx) error {
	id := c.Params("id")

	res, err := DB.Exec(
		"DELETE FROM contacts WHERE id=$1", id,
	)

	if err != nil {
		return c.Status(500).SendString("Delete error")
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}

	return c.SendStatus(204)
}
