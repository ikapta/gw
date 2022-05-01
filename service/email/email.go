package email

import (
	"fmt"

	"kapta.gw/config"
	"kapta.gw/utils"
	"github.com/go-gomail/gomail"
)

type MsgConf struct {
    To          []string
    Subject     string
    Body        string
    AliasFrom   string
}

func Send(cfg MsgConf) error {
  mailConf, err := getEmailConfig()

  if err != nil {
      return err
  }

  if len(cfg.To) <= 0 {
    return fmt.Errorf("email config error: To is empty")
  }

  m := gomail.NewMessage()

  fmt.Printf(fmt.Sprintf("%s <%s>", cfg.AliasFrom, mailConf.SMTP_EMAIL))

  m.SetHeader("To", utils.DeduplicateStr(cfg.To)...)
  m.SetHeader("Subject", cfg.Subject)
  m.SetBody("text/html", cfg.Body)
  m.SetAddressHeader("From", mailConf.SMTP_EMAIL, cfg.AliasFrom)
  // m.SetHeader("Cc", "")
  // m.Attach("")

  d := gomail.NewPlainDialer(mailConf.SMTP_HOST, mailConf.SMTP_PORT, mailConf.SMTP_EMAIL, mailConf.SMTP_PASS)

  if err = d.DialAndSend(m); err != nil {
    fmt.Printf("send email error: %v", err)
    return fmt.Errorf("send email error: %v", err)
  }

  return nil
}

func getEmailConfig() (*config.EmailConf, error) {
  _conf := config.GetAll().Email

  if len(_conf.SMTP_EMAIL) <= 0 {
      return nil, fmt.Errorf("email config error: SMTP_EMAIL is empty")
  }
  return _conf, nil
}