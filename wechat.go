package wechat

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	endpoint        = "https://api.weixin.qq.com"
	loginPath       = "/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	accessTokenPath = "/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	jsonContentType = "application/json"
)

type Wechat struct {
	*http.Client

	appid     string
	secretKey string

	accessToken          string
	accessTokenExpiresAt int64
}

func NewWechat(cfg *Config) *Wechat {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		Timeout: cfg.TimeoutDuration(),
	}

	return &Wechat{
		Client:    client,
		appid:     cfg.AppID,
		secretKey: cfg.SecretKey,
	}
}

func (wechat *Wechat) Login(code string) (output *LoginOutput, err error) {
	data, err := wechat.doReqeust(buildAPI(loginPath, wechat.appid, wechat.secretKey, code))
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		return
	}

	if output.WechatError != nil {
		return nil, output.Error()
	}

	return
}

func (wechat *Wechat) AccessToken() (token string, err error) {
	if wechat.accessToken != "" && time.Now().Unix() < wechat.accessTokenExpiresAt {
		return wechat.accessToken, nil
	}

	data, err := wechat.doReqeust(buildAPI(accessTokenPath, wechat.appid, wechat.secretKey))
	if err != nil {
		return
	}

	var output AccessTokenOutput
	err = json.Unmarshal(data, &output)
	if err != nil {
		return
	}

	if output.WechatError != nil {
		return token, output.Error()
	}

	wechat.accessToken = output.AccessToken
	wechat.accessTokenExpiresAt = time.Now().Unix() + output.ExpiresIn

	return output.AccessToken, nil
}

func (wechat *Wechat) doReqeust(url string) ([]byte, error) {
	resp, err := wechat.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func buildAPI(api string, params ...interface{}) string {
	return fmt.Sprintf(endpoint+api, params...)
}
