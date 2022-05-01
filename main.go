package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/recover"
  "github.com/gofiber/fiber/v2/middleware/requestid"

  "kapta.gw/config"
  "kapta.gw/router"
)

func main() {

  conf := config.GetAll()

  app := fiber.New(fiber.Config{
    AppName: conf.App.NAME,
  })

  app.Config()

  app.Use(recover.New())

  app.Use(requestid.New())

  app.Use(cors.New(cors.Config{
    AllowOrigins: strings.Join(conf.App.ALLOW_ORIGINS, ","),
    AllowCredentials: true,
  }))

  app.Static("/", "./public")

  router.Setup(app)

  app.Listen(conf.App.PORT)
}
