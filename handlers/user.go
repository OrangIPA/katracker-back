package handlers

import (
	"database/sql"
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

func NewUser(c *fiber.Ctx) error {
	reqBody := newPersonParam{}

	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	db.Conn.MustExec("INSERT INTO person(username, pass) VALUES ($1, $2);", reqBody.Username, reqBody.Password)

	return nil
}

func AllUser(c *fiber.Ctx) error {
	users := []person{}
	err := db.Conn.Select(&users, "SELECT * FROM person;")
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	user := person{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("invalid id")
	}

	err = db.Conn.Get(&user, "SELECT * FROM person WHERE id=$1", id)
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

	_, err = db.Conn.Exec("UPDATE person SET pass=$1 WHERE id=$2", reqBody.Password, id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
