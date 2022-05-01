package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
  WebhookFormat = "https://open.feishu.cn/open-apis/bot/v2/hook/%s"
)

type Language string

const (
	LangChinese  Language = "zh_cn"
	LangEnglish  Language = "en_us"
	LangJapanese Language = "ja_jp"
)

// 实现@功能 https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#e1cdee9f
func AtAllInPost() string {
	return `<at user_id="all"></at>`
}

func AtUserInPost(id string) string {
	return fmt.Sprintf(`<at user_id="%s"></at>`, id)
}

// msg_type = interactive
func AtAllInCard() string {
  return `<at id=all></at>`
}

func AtUserInCard(id string) string {
	return fmt.Sprintf(`<at id="%s"></at>`, id)
}


func AtUserByEmail(email string) string {
  return fmt.Sprintf(`<at user_email="%s"></at>`, email)
}

func AtUserName(id string, name ...string) string {
	var s string
	if len(name) != 0 {
		s = name[0]
	}
	return fmt.Sprintf(`<at user_id="%s">%s</at>`, id, s)
}

// 仅支持部分
// 语法详情: https://open.feishu.cn/document/ukTMukTMukTM/uADOwUjLwgDM14CM4ATN

func Italics(s string) string {
	return fmt.Sprintf("*%s*", s)
}

func Bold(s string) string {
	return fmt.Sprintf("**%s**", s)
}

func Strikethrough(s string) string {
	return fmt.Sprintf("~~%s~~", s)
}

func Link(url string) string {
	return fmt.Sprintf("<a>%s</a>", url)
}

func TextLink(text, url string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

func Image(hoverText, imageKey string) string {
	return "!" + TextLink(hoverText, imageKey)
}

func HorizontalRule() string {
	return ` ---`
}


// 签名校验
//  https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#348211be
func genSign(secret string, timestamp int64) (string, error) {
	sign := fmt.Sprintf("%d\n%s", timestamp, secret)

	var data []byte
	h := hmac.New(sha256.New, []byte(sign))
	if _, err := h.Write(data); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func execPost(webhook string, body []byte) error {
  agent := fiber.AcquireAgent()
  res := fiber.AcquireResponse()
  req := agent.Request()

  defer fiber.ReleaseAgent(agent)
  defer fiber.ReleaseResponse(res)

  req.Header.SetMethod(fiber.MethodPost)
  req.SetRequestURI(webhook)
  req.SetBody([]byte(body))

  if err := agent.Parse(); err != nil {
    return err
  }

	if err := agent.HostClient.Do(req, res); err != nil {
    return err
	}

  return nil
}