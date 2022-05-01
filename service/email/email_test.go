package email_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"kapta.gw/service/email"
)

func TestSend(t *testing.T) {
  var (
    to []string = []string{"kapta.fu@outlook.com"}
    subject = "[发布提醒] 测试邮件"
    htmlBody = "<a href='https://kaifa.baidu.com'>kaifa</a>"
    aliasFrom = "FE"
  )

  err := email.Send(email.MsgConf{
    To: to,
    Subject: subject,
    Body: htmlBody,
  	AliasFrom: aliasFrom,
  })

  if err != nil {
    t.Error(err)
  } else {
    t.Log("send email success")
  }
}

//获取程序执行目录
func TestGetRunPath(t *testing.T)  {
  path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
  fmt.Print(path)
}

//获取程序执行目录
func TestGetRunPath2(t *testing.T) {
  file, _ := exec.LookPath(os.Args[0])
  path, _ := filepath.Abs(file)
  index := strings.LastIndex(path, string(os.PathSeparator))
  ret := path[:index]
  fmt.Print(ret)
}