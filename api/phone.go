package api

import (
	"errors"
	"fmt"

	"github.com/songjiayang/wechat/util"
)

type PhoneNumberResult struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	Watermark       *Watermark `json:"watermark"`
}

type Watermark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

func GetPhoneNumber(appID, sessionKey, iv, encryptedData string) (ret PhoneNumberResult, err error) {
	if err = util.DecryptData(sessionKey, iv, encryptedData, &ret); err != nil {
		return
	}

	if ret.Watermark.AppID != appID {
		err = errors.New(fmt.Sprintf("appid not matched: %s != %s", appID, ret.Watermark.AppID))
	}

	return
}
