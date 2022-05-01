package bot

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Bot interface {
  SendText(content string) error
  SendPost(p Post, ps ...Post) error
  SendImage(imageKey string) error
  SendCard(bgColor CardTitleBgColor, cfg CardConfig, c Card, more ...Card) error
}

type bot struct {
	webhook string
	secretKey string
}

type BotOption func(*bot)

func NewBot(wbhk string, opts ...BotOption) Bot {
  b := new(bot)

  wbhk = strings.TrimSpace(wbhk)

	if !strings.Contains(wbhk, "open.feishu.cn") {
    b.webhook = fmt.Sprintf(WebhookFormat, wbhk)
  } else {
    b.webhook = wbhk
  }
	for _, fn := range opts {
		fn(b)
	}
	return b
}

func WithSecretKey(key string) BotOption {
	return func(b *bot) {
		b.secretKey = strings.TrimSpace(key)
	}
}

func (b *bot) SendText(content string) error {
	return b.send(NewText(content))
}

func (b *bot) SendImage(imageKey string) error {
  return b.send(NewImage(imageKey))
}

func (b *bot) SendCard(bgColor CardTitleBgColor, cfg CardConfig, c Card, more ...Card) error {
  return b.send(NewCard(bgColor, cfg, c, more...))
}

func (b *bot) SendPost(p Post, ps ...Post) error {
	return b.send(NewPost(p, ps...))
}

func (b *bot) send(msg map[string]interface{}) (err error) {
	if b.secretKey != "" {
		ts := time.Now().Unix()
		signed, err := genSign(b.secretKey, ts)
		if err != nil {
			return err
		}
		msg["timestamp"] = ts
		msg["sign"] = signed
	}

	var msgBody []byte
	if msgBody, err = json.Marshal(msg); err != nil {
		return err
	}

  err = execPost(b.webhook, msgBody)
	return err
}
