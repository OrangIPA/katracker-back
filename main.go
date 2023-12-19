package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgres://postgres:semangatpagi@localhost:5432/test?sslmode=disable")

	app := fiber.New()

	route(app)

	log.Fatal(app.Listen(":3000"))
}
