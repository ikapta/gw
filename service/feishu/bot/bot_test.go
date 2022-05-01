package bot_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"kapta.gw/config"
	"kapta.gw/service/feishu/bot"
)

var testWh = config.GetAll().Feishu.WBHK_XIAOTAN
var xiaotan = bot.NewBot(testWh)

func TestPushText(t *testing.T) {
  xiaotan.SendText("Hello world." + bot.Strikethrough("abc"))
}

func TestPushTextAtAll(t *testing.T) {
  xiaotan.SendText(fmt.Sprintf("Hello world. %s", bot.AtAllInPost()))
}

func TestSendCardDev(t *testing.T) {
  xiaotan.SendCard(
    bot.BgColorBlue, nil,
    bot.WithCard(
      bot.LangChinese,
      "[DEV]发布提醒",
      bot.WithCardElementMarkdown("测试环境发布成功。👉[点击去测试](https://app.cew-dev.com/)"),
    ),
  )
}

func TestSendCardProd(t *testing.T) {
  xiaotan.SendCard(
    bot.BgColorGreen, nil,
    bot.WithCard(
      bot.LangChinese,
      "[PROD]发布提醒",
      bot.WithCardElementMarkdown(fmt.Sprintf("生产环境发布成功。%s", bot.AtAllInCard())),
      bot.WithCardElementMarkdown("👉[点击去测试](https://app.carbonnt.com/)"),
    ),
  )
}

func TestStr2Map(t *testing.T) {
  str := `{"data":{"base":"BTC","currency":"USD","amount":4225.87}}`

  data := make(map[string]interface{})
  err := json.Unmarshal([]byte(str), &data)

  if err != nil {
    t.Error(err)
  } else {
    fmt.Printf("%s", data["data"])
    t.Log(data)
  }
}

func TestJSON2Str(t *testing.T) {
  var data = map[string]interface{}{
    "msg_type": "text",
    "content": "Hello world!",
  }

  str, err := json.Marshal(data)

  if err != nil {
    t.Error(err)
  } else {
    fmt.Println(str)
  }
}


