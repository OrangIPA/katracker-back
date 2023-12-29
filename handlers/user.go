package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"strconv"

	"github.com/OrangIPA/katracker-back/db"
	"github.com/gofiber/fiber/v2"
)

type newPersonParam struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type person struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Pass     string `json:"password" db:"pass"`
}

type getPerson struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

func NewUser(c *fiber.Ctx) error {
	reqBody := newPersonParam{}

	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	digest := sha256.Sum256([]byte(reqBody.Password))

	db.Conn.MustExec("INSERT INTO person(username, pass) VALUES ($1, $2);", reqBody.Username, hex.EncodeToString(digest[:]))

	return nil
}

func AllUser(c *fiber.Ctx) error {
	users := []getPerson{}
	err := db.Conn.Select(&users, "SELECT id, username FROM person;")
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	user := getPerson{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("invalid id")
	}

	err = db.Conn.Get(&user, "SELECT id, username FROM person WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return c.SendStatus(fiber.StatusNotFound)
	} else if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(user)
}

func DelUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("invalid id")
	}

	_, err = db.Conn.Exec("DELETE FROM person WHERE id = $1", id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateUser(c *fiber.Ctx) error {
	reqBody := newPersonParam{}
	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("invalid id")
	}

	digest := sha256.Sum256([]byte(reqBody.Password))

	_, err = db.Conn.Exec("UPDATE person SET username=$1, pass=$2 WHERE id=$3", reqBody.Username, hex.EncodeToString(digest[:]), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func ChangeUsername(c *fiber.Ctx) error {
	reqBody := newPersonParam{}
	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("invalid id")
	}

	_, err = db.Conn.Exec("UPDATE person SET username=$1 WHERE id=$2", reqBody.Username, id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func ChangePassword(c *fiber.Ctx) error {
	reqBody := newPersonParam{}

	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("invalid id")
	}

	digest := sha256.Sum256([]byte(reqBody.Password))

	_, err = db.Conn.Exec("UPDATE person SET pass=$1 WHERE id=$2", hex.EncodeToString(digest[:]), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
