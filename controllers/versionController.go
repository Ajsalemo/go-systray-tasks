package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type VersionControllerMessage struct {
	Msg string
  }

func VersionController(c *fiber.Ctx) error {
	res := VersionControllerMessage{
		Msg: time.Now().Format("2006/01/02") + " | v1.0",
	}
	return c.JSON(res)
}