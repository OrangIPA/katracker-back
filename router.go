package main

import (
	"github.com/OrangIPA/katracker-back/handlers"
	"github.com/gofiber/fiber/v2"
)

func route(app *fiber.App) {
	app.Post("/person", handlers.NewUser)
	app.Get("/person", handlers.AllUser)
	app.Get("/person/:id", handlers.GetUser)
	app.Delete("/person/:id", handlers.DelUser)
	app.Put("/person/username/:id", handlers.ChangeUsername)
	app.Put("/person/password/:id", handlers.ChangePassword)
}
