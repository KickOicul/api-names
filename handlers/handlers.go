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

func SugiereNombreGato(c *fiber.Ctx) error {
	// Parsea el cuerpo JSON de la solicitud
	var requestBody map[string]string
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al analizar la solicitud"})
	}

	// Obtiene el color de la solicitud
	color, ok := requestBody["color"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Se requiere el campo 'color' en la solicitud"})
	}

	// Obtiene el nombre sugerido de la solicitud
	suggestedName, ok := requestBody["suggestedName"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Se requiere el campo 'suggestedName' en la solicitud"})
	}

	// Agrega el nombre sugerido a la lista de nombres para el color específico
	gatos.AddSuggestedCatName(color, suggestedName)

	return c.JSON(fiber.Map{"message": "Sugerencia de nombre de gato recibida con éxito"})
}
