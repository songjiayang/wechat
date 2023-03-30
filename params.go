package wechat

import (
	"errors"
	"fmt"
)

type LoginOutput struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`

	WechatError
}

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
type PhoneNumberOutput struct {
	Info PhoneInfo `json:"phone_info"`
	WechatError
}

func (p *PhoneNumberOutput) PhoneNumber() string {
	if p.Info.PurePhoneNumber != "" {
		return p.Info.PurePhoneNumber
	}

	return p.Info.PhoneNumber
}

type PhoneInfo struct {
	CountryCode     string `json:"countryCode"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
}

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183
type AccessTokenOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`

	WechatError
}

type WechatError struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (e *WechatError) Error() error {
	if e.Code == 0 {
		return nil
	}

	return errors.New(fmt.Sprintf("errcode: %d, errmsg: %s", e.Code, e.Msg))
}
