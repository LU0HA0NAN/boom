package warn

import (
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

var WebHookAddr = ""

func WarnToWechat(msg string) {
	if WebHookAddr == "" {
		return
	}

	body := `
		{
			"msgtype": "markdown",
			"markdown": {
				"content": "CONTENT"
			}
		}
		`
	body = strings.ReplaceAll(body, "CONTENT", msg)
	_, _ = resty.New().R().SetHeader("Content-Type", "application/json").SetBody(body).Post(WebHookAddr)
}

func WarnToWechatII(webhook, msg string) {
	if webhook == "" {
		return
	}

	body := `
		{
			"msgtype": "markdown",
			"markdown": {
				"content": "CONTENT"
			}
		}
		`
	body = strings.ReplaceAll(body, "CONTENT", msg)
	_, _ = resty.New().R().SetHeader("Content-Type", "application/json").SetBody(body).Post(webhook)
}

func init() {
	WebHookAddr = os.Getenv("wechat_webhook")
}
