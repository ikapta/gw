package controller

import (
	"fmt"
	"io/ioutil"

	"kapta.gw/service/deploy"
	"github.com/gofiber/fiber/v2"
)

func DeployController(c *fiber.Ctx) error {
  return c.SendFile("public/deploy.html")
}

func DeployProjLogController(c *fiber.Ctx) error {
  var logName = c.Params("name")
  fileData, err := ioutil.ReadFile("./log/" + logName)

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error: %v\n", err.Error()))
  }

  return c.JSON(fiber.Map{
    "log": string(fileData),
    // "deployTotal": deployLimit.Total, // request deploy total
    // "deployWait": deployLimit.Total - deployLimit.MaxQ, // request deploy wait count
  })
}

func DeployProjController(c *fiber.Ctx) error {
  var isDryRun = c.Query("dryRun") != ""

  var deployProd = c.Query("deployEnv") == "prod"

  var projName = c.Params("name")

  var requestId = c.Context().Response.Header.Peek(fiber.HeaderXRequestID)

  var logFileName = fmt.Sprintf("deploy-%s-%s.log", projName, requestId)

  go deploy.ImpDeployProcess.BuildAndDeploy(&deploy.DeployProcessType{
    ProjName: projName,
    LogFileName: logFileName,
    IsDryRun: isDryRun,
    DeployProd: deployProd,
  })

  return c.SendString(logFileName)
}
