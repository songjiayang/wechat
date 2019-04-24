package api

import (
	"github.com/golib/assert"
	"os"
	"testing"
)

func TestGetPhoneNumber(t *testing.T) {
	assertion := assert.New(t)

	sessionKey := os.Getenv("WECHAT_SESSION_KEY")
	iv := os.Getenv("WECHAT_IV")
	encryptedData := os.Getenv("WECHAT_ENCRYPTED_DATA")

	ret, err := GetPhoneNumber(os.Getenv("WECHAT_APPID"), sessionKey, iv, encryptedData)

	assertion.Nil(err)
	assertion.NotEmpty(ret.PhoneNumber)
}
