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
)

type Wechat struct {
	*http.Client

	appID     string
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
		appID:     cfg.AppID,
		secretKey: cfg.SecretKey,
	}
}

func (w *Wechat) Login(code string) (output *LoginOutput, err error) {
	data, err := w.doRequest(buildAPI(loginPath, w.appID, w.secretKey, code))
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

func (w *Wechat) AccessToken() (token string, err error) {
	if w.accessToken != "" && time.Now().Unix() < w.accessTokenExpiresAt {
		return w.accessToken, nil
	}

	data, err := w.doRequest(buildAPI(accessTokenPath, w.appID, w.secretKey))
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

	w.accessToken = output.AccessToken
	w.accessTokenExpiresAt = time.Now().Unix() + output.ExpiresIn

	return output.AccessToken, nil
}

func (w *Wechat) doRequest(url string) ([]byte, error) {
	resp, err := w.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func buildAPI(api string, params ...interface{}) string {
	return fmt.Sprintf(endpoint+api, params...)
}
