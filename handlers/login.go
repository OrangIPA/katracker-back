package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"

	"github.com/OrangIPA/katracker-back/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type user struct {
	Id       int    `db:"id"`
	Username string `json:"username" form:"username" xml:"username" db:"username"`
	Password string `json:"password" form:"password" xml:"password" db:"pass"`
}

func Login(c *fiber.Ctx) error {
	loginParam := user{}

	if err := c.BodyParser(&loginParam); err != nil {
		return err
	}

	user := user{}
	if err := db.Conn.Get(&user, "SELECT id, username, pass FROM person WHERE username = $1", loginParam.Username); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.SendStatus(404)
		}
		return err
	}

	digest := sha256.Sum256([]byte(loginParam.Password))

	if hex.EncodeToString(digest[:]) != user.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"sub":     user.Username,
		"iat":     time.Now(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"user_id": user.Id,
	}

	secret := os.Getenv("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
