package controllers

import (
    "github.com/gofiber/fiber/v2"
)

type EnvControllerMessage struct {
	TitleMsg string
	AvailableEnvVar map[string]string
  }

func EnvController(c *fiber.Ctx) error {
	res := EnvControllerMessage{
		TitleMsg: "Available environment variables that can be configured",
		AvailableEnvVar: map[string]string{
			"DISABLE_BACKLOG_TITLE": "If DISABLE_BACKLOG_TITLE is 'true' / 'TRUE' or 1 - this disables the hotkey for backlog titles",
			"DISABLE_BACKLOG_BODY": "If DISABLE_BACKLOG_BODY is 'true' / 'TRUE' or 1 - this disables the hotkey for backlog bodies",
			"DISABLE_AGED_TITLE": "If DISABLE_AGED_TITLE is 'true' / 'TRUE' or 1 - this disables the hotkey for aged titles",
			"DISABLE_AGED_BODY": "If DISABLE_AGED_BODY is 'true' / 'TRUE' or 1 - this disables the hotkey for aged bodies",
			"DISABLE_FDR_TITLE": "If DISABLE_FDR_TITLE is 'true' / 'TRUE' or 1 - this disables the hotkey for FDR titles",
			"DISABLE_FDR_BODY": "If DISABLE_FDR_BODY is 'true' / 'TRUE' or 1 - this disables the hotkey for FDR bodies",
			"PORT": "The port number to run the application on - this defaults to 3080 unless otherwise set",
		},

	}
	return c.JSON(res)
}