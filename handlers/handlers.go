package handlers

import (
	"api-gatos/utilities/barcos"
	"api-gatos/utilities/gatos"

	"github.com/gofiber/fiber/v2"
)

func GetCatName(c *fiber.Ctx) error {
	color := c.Params("color")

	if color == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Se requiere el parámetro color"})
	}

	foto, err := gatos.GetCatNameByColor(color)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"nombre": foto})
}

func GeneraNombreBarco(c *fiber.Ctx) error {
	shipName := barcos.GeneraNombreBarcoAleatorio()
	return c.JSON(fiber.Map{"shipName": shipName})
}
