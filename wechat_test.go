package wechat

import (
	"os"
	"testing"

	"github.com/golib/assert"
)

var wClient *Wechat

func TestMain(m *testing.M) {
	wClient = NewWechat(&Config{
		os.Getenv("WECHAT_APPID"),
		os.Getenv("WECHAT_SECRETKEY"),
		30,
	})

	os.Exit(m.Run())
}

func TestAccessToken(t *testing.T) {
	assertion := assert.New(t)
	token, err := wClient.AccessToken()

	assertion.Nil(err)
	assertion.NotEmpty(token)
}
