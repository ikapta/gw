package controller

import (
	"github.com/gofiber/fiber/v2"
)

func HomeController(c *fiber.Ctx) error {
  return c.SendFile("public/index.html")
}