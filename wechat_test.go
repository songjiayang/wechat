package wechat

import (
	"os"
	"testing"

	"github.com/golib/assert"
)

func TestAccessToken(t *testing.T) {
	assertion := assert.New(t)
	wClient := NewWechat(&Config{
		os.Getenv("WECHAT_APPID"),
		os.Getenv("WECHAT_SECRETKEY"),
		30,
	})

	token, err := wClient.AccessToken()

	assertion.Nil(err)
	assertion.NotEmpty(token)
}
