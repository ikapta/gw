package config

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"kapta.gw/utils"
	"github.com/spf13/viper"
)

type Configs struct {
  App *AppConf `yaml:"app"`
  Deploy *DeployConf `yaml:"deploy"`
  Email *EmailConf `yaml:"email"`
  SSH *SSHConf `yaml:"ssh"`
  Feishu *FeishuConf `yaml:"feishu"`
}

type AppConf struct {
  NAME string `yaml:"name"`
  PORT string `yaml:"port"`
  ALLOW_ORIGINS []string `yaml:"allow_origins"`
}

type DeployConf struct {
  PROJ_WORKSPACE string `yaml:"proj_workspace"`
  RUNTIME_MODE string `yaml:"RUNTIME_MODE"` // deploy mode: IO_SSH | Docker | IO_IN_SERVER
}

type EmailConf struct {
  SMTP_HOST string `yaml:"smtp_host"`
  SMTP_PORT int `yaml:"smtp_port"`
  SMTP_EMAIL string `yaml:"smtp_email"`
  SMTP_PASS string `yaml:"smtp_pass"`
}

type SSHConf struct {
  HOST string `yaml:"host"`
  NAME string `yaml:"name"`
  PASS string `yaml:"pass"`
}

type FeishuConf struct {
  WBHK_XIAOTAN string `yaml:"wbhk_xiaotan"`
}

var V *viper.Viper
var confInst Configs

func GetAll() Configs {
  if V == nil {
    V = initConfig()
  }

  if err := V.Unmarshal(&confInst); err != nil {
    fmt.Println(err.Error())
  }

  fmt.Printf("confInst: %#v\n", confInst)
  return confInst
}

// get from default and merge local env config.
func initConfig() (*viper.Viper) {
  vInst := viper.New()

  absPath := utils.GetRootPath()

  vInst.SetConfigFile(absPath + "/config/conf.default.yml")

  configBytes, err := ioutil.ReadFile(absPath + "/config/conf.local.yml")

  if err != nil {
      fmt.Printf("no `conf.local.yml` file merge to default: %s \n", err)
  } else {
    vInst.MergeConfig(bytes.NewBuffer(configBytes))
  }

  return vInst
}

