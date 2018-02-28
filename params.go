package wechat

import (
	"errors"
	"fmt"
)

type LoginOutput struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`

	*WechatError
}

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
type AccessTokenOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`

	*WechatError
}

type WechatError struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (e *WechatError) Error() error {
	return errors.New(fmt.Sprintf("errcode: %d, errmsg: %s", e.Code, e.Msg))
}
