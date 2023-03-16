package wechat

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	token, err := wClient.AccessToken()

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
