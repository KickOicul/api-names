package routes

import (
	"api-gatos/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/sugiereNombreGato", handlers.SugiereNombreGato)
	app.Get("/getCatName/:color", handlers.GetCatName)
	app.Get("/generaNombreBarco", handlers.GeneraNombreBarco)
}
