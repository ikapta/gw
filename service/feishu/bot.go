package feishu

import (
	"fmt"

	"kapta.gw/config"
	"github.com/gofiber/fiber/v2"
)

const (

)

type MsgBody[T any] struct {
  msg_type string
  content T
}

type MsgBodyText struct {
  msg_type string
  text string
}

// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#c9e08671

// 如果需要换行，需要转义: \\n
func SendText(msg string) error {

//    buf := bytes.NewBufferString("新更新提醒\n")
// 		buf.WriteString("🤓所有人👉" + fsBotAPI.StrMentionAll() + "\n")
// 		buf.WriteString("🤔你是谁👉" + fsBotAPI.StrMentionByOpenID("ou_c99c5f35d542efc7ee492afe11af19ef") + "\n")


  err := Send(
    "text",
    fmt.Sprintf(" {\"text\":\"%s\"} ", msg),
  )
  return err
}

func Send(msgType string, content interface{}) error {
  webhook := config.GetAll().Feishu.WBHK_XIAOTAN

  if len(webhook) <= 0 {
    return fmt.Errorf("feishu config error: WBHK_XIAOTAN is empty")
  }

  agent := fiber.AcquireAgent()
  res := fiber.AcquireResponse()
  req := agent.Request()

  defer fiber.ReleaseAgent(agent)
  defer fiber.ReleaseResponse(res)

  req.Header.SetMethod(fiber.MethodPost)
  req.SetRequestURI(webhook)
  req.SetBody([]byte(fmt.Sprintf("{\"msg_type\":\"%s\",\"content\":%s}", msgType, content)))

  if err := agent.Parse(); err != nil {
    return err
  }

	if err := agent.HostClient.Do(req, res); err != nil {
    return err
	}

  return nil
}