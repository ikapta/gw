package router

import (
  "github.com/gofiber/fiber/v2"
  "kapta.gw/controller"
)

func Setup(app *fiber.App) {
  app.Get("/", controller.HomeController)

  app.Post("/gw/*", controller.GwController)

  deploy := app.Group("/deploy")
  deploy.Get("/", controller.DeployController)
  deploy.Post("/proj/:name", controller.DeployProjController)
  deploy.Get("/proj/log/:name", controller.DeployProjLogController)
}