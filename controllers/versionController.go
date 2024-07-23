package controllers

import (
    "github.com/gofiber/fiber/v2"
)

type VersionControllerMessage struct {
	Msg string
  }

func VersionController(c *fiber.Ctx) error {
	res := VersionControllerMessage{
		Msg: "e73fb59d1220ab4610183521507693c4b5e53c69 | v1.0",
	}
	return c.JSON(res)
}