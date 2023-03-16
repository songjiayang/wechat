package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPhoneNumber(t *testing.T) {

	if os.Getenv("WECHAT_SESSION_KEY") == "" {
		return
	}

	sessionKey := os.Getenv("WECHAT_SESSION_KEY")
	iv := os.Getenv("WECHAT_IV")
	encryptedData := os.Getenv("WECHAT_ENCRYPTED_DATA")

	ret, err := GetPhoneNumber(os.Getenv("WECHAT_APPID"), sessionKey, iv, encryptedData)

	assert.Nil(t, err)
	assert.NotEmpty(t, ret.PhoneNumber)
}
