package main

import (
	"log"
	"os"

	"github.com/OrangIPA/katracker-back/db"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
}

func main() {
	app := fiber.New()

	route(app)

	log.Fatal(app.Listen(os.Getenv("host")))
}
