package deploy

import (
	"errors"
	"fmt"

	"kapta.gw/config"
	"kapta.gw/service/email"
	"kapta.gw/service/feishu/bot"
	"kapta.gw/service/ssh"
	"kapta.gw/utils"
)

type deployService struct {}

var ImpDeploy = &deployService{}

type ProjInfoModel struct {
  Name string
  GitPath string
  PubSerPath string // the value for deploy to serve by env.
  pubDevServerPath string
  pubProdServerPath string
  PubServerZipPath string
  DevDomain string
  ProdUrl string
}

var a = &ProjInfoModel{
}

var b = &ProjInfoModel{
}

var projInfoMap = map[string]*ProjInfoModel{
  "a" : a,
  "b" : b,
}

func (p *deployService) GetProjInfoByName(name string, deployProd bool) (*ProjInfoModel, error) {
  proj := projInfoMap[name]

  if proj == nil {
    return nil, errors.New("proj not exist - " + name)
  }

  // default deploy to dev
  if deployProd {
    proj.PubSerPath = proj.pubProdServerPath
  } else {
    proj.PubSerPath = proj.pubDevServerPath
  }

  return proj, nil
}

func (p *deployService) GenZipName(projName string) string {
  return fmt.Sprintf("%s_%s", projName, "develop.tar.gz")
}

// 部署成功发送通知邮件
func (p *deployService) SendDeployMsg(projInfo *ProjInfoModel, abcConf AbcConfig, projWorkPath string, deployProd bool) {
  var (
    subj=            ""
    defaultNotices=  []string{"kapta.fu@outlook.com"}
    mergedNotices=   append(abcConf.Notices, defaultNotices...)
  )

  fmt.Printf("send deploy msg to: %#v\n", mergedNotices)

  if deployProd {
    subj = fmt.Sprintf("[发布提醒] 生产环境发布成功 %s", projInfo.Name)
  } else {
    subj = fmt.Sprintf("[发布提醒] 测试环境发布成功 %s", projInfo.Name)
  }

  email.Send(email.MsgConf{
    To:       mergedNotices,
    Subject:  subj,
    AliasFrom: "FE",
    Body:     fmt.Sprintf(`
    <div>
      您的前端项目 <a href="%s">%s</a>，成功部署到测试环境。<br/>
      <a href="%s">点击查看</a><br/><br/><br/><br/><br/>
      <small>为什么会收到这封邮件，因为您是该项目的开发者。</small>
    </div>
    `, projInfo.GitPath, projInfo.Name, projInfo.DevDomain),
  })
}

func (p *deployService) SendDeployCardMsgForFeishu(projInfo *ProjInfoModel, deployProd bool) {
  deployWebhook := config.GetAll().Feishu.WBHK_XIAOTAN

  if len(deployWebhook) <= 0 {
    fmt.Printf("no deploy webhook, skip send deploy feishu card msg\n")
    return
  }

  xiaotan := bot.NewBot(deployWebhook)

  if deployProd {
    xiaotan.SendCard(
      bot.BgColorGreen, nil,
      bot.WithCard(
        bot.LangChinese,
        "[PROD]发布提醒",
        bot.WithCardElementMarkdown(fmt.Sprintf("%s生产环境发布成功。%s", projInfo.Name, bot.AtAllInCard())),
        bot.WithCardElementMarkdown(fmt.Sprintf("👉[点击去测试](%s)", projInfo.ProdUrl)),
      ),
    )
  } else {
    xiaotan.SendCard(
      bot.BgColorBlue, nil,
      bot.WithCard(
        bot.LangChinese,
        "[DEV]发布提醒",
        bot.WithCardElementMarkdown(
          fmt.Sprintf("%s测试环境发布成功。👉[点击去测试](%s)", projInfo.Name, projInfo.DevDomain),
        ),
      ),
    )
  }
}

type AbcConfig struct {
  Notices []string `yaml:"notices"`
}

// 读取前端项目根路径的自定义配置abc.yml
func (p *deployService) GetAbcConf(projPath string) (AbcConfig, error) {
  abcPath := fmt.Sprintf("%s/abc.yml", projPath)
  abc, _, err := utils.ReadConfFile[AbcConfig](abcPath)

  if err != nil {
    return AbcConfig{}, errors.New("read abc.yml error - " + err.Error())
  }

  return abc, nil
}

func (p *deployService) RunCmdInServer(cmd string) (string, error) {
  cfg := config.GetAll().SSH

  return ssh.Once(
    ssh.SSHConf{
      HOST: cfg.HOST,
      NAME: cfg.NAME,
      PASS: cfg.PASS,
    },
    cmd)
}

// todo: count https://stackoverflow.com/questions/70046636/how-to-check-if-sync-waitgroup-done-is-called-in-unit-test
// type DeployLimit struct {
//   ch chan int
//   Total int
//   ReqId string
//   WaitNum int
//   MaxQ int
// }

// func NewDeployLimit() DeployLimit {
//   max := 2 // max concurrence only 2
//   return DeployLimit{ch: make(chan int, max), MaxQ: max}
// }

// func (g *DeployLimit) Add() {
//   g.Total++
//   g.ch <- 1
// }

// func (g *DeployLimit) Done() {
//   g.Total--
//   <-g.ch
// }