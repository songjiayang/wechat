package wechat

import (
	"encoding/json"
	"testing"

	"github.com/golib/assert"
)

func TestAccessTokenOutput(t *testing.T) {
	assertion := assert.New(t)

	jsonStr := `{"errcode": 40013, "errmsg":"invalid appid"}`
	var output AccessTokenOutput
	err := json.Unmarshal([]byte(jsonStr), &output)

	assertion.Nil(err)
	assertion.NotNil(output.WechatError)
	assertion.NotNil(output.Error())

	jsonStr = `{"access_token": "xxx", "expires_in": 123}`
	var output2 AccessTokenOutput
	err = json.Unmarshal([]byte(jsonStr), &output2)

	assertion.Nil(err)
	assertion.Nil(output2.WechatError)
	assertion.Equal(output2.AccessToken, "xxx")
	assertion.Equal(output2.ExpiresIn, int64(123))
}
