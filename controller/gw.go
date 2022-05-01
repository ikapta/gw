package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
  "kapta.gw/config"
)

/**
 * todo: use proxy Balancer creates a load balancer among multiple upstream servers
 * https://github.com/gofiber/fiber/tree/master/middleware/proxy
 */
func GwController(c *fiber.Ctx) error {
  apiPath := c.Path()[len(c.Route().Path) - 1:]

  if len(apiPath) == 0  {
    return c.Status(fiber.StatusBadRequest).SendString("Error: missing api path!")
  }

  agent := fiber.AcquireAgent()
  res := fiber.AcquireResponse()
  req := agent.Request()

  defer fiber.ReleaseAgent(agent)
  defer fiber.ReleaseResponse(res)

  c.Context().Request.CopyTo(req)

  req.SetRequestURI(fmt.Sprintf("%s/%s", config.GetAll().App.GW_API, apiPath))
  req.SetBody(c.Body())

  if err := agent.Parse(); err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Parse error: %v\n", err))
  }

	if err := agent.HostClient.Do(req, res); err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Do error: %v\n", err))
	}

  c.Context().Response.Header.Set("content-encoding", "gzip")

  return c.Status(res.StatusCode()).Send(res.Body())
}