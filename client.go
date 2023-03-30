package wechat

import (
	"bytes"
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
	getPhonePath    = "/wxa/business/getuserphonenumber?access_token=%s"
	accessTokenPath = "/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type Client struct {
	*http.Client

	appID     string
	secretKey string

	accessToken          string
	accessTokenExpiresAt int64
}

func NewClient(cfg *Config) *Client {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		Timeout: cfg.TimeoutDuration(),
	}

	return &Client{
		Client:    client,
		appID:     cfg.AppID,
		secretKey: cfg.SecretKey,
	}
}

func (c *Client) AccessToken() (token string, err error) {
	if c.accessToken != "" && time.Now().Unix() < c.accessTokenExpiresAt {
		return c.accessToken, nil
	}

	data, err := c.doRequest(buildAPI(accessTokenPath, c.appID, c.secretKey))
	if err != nil {
		return
	}

	var output AccessTokenOutput
	if err = json.Unmarshal(data, &output); err != nil {
		return
	}

	if output.Error() != nil {
		return token, output.Error()
	}

	c.accessToken = output.AccessToken
	c.accessTokenExpiresAt = time.Now().Unix() + output.ExpiresIn

	return output.AccessToken, nil
}

func (c *Client) Login(code string) (output *LoginOutput, err error) {
	data, err := c.doRequest(buildAPI(loginPath, c.appID, c.secretKey, code))
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		return
	}

	if output.Error() != nil {
		return nil, output.Error()
	}

	return
}

func (c *Client) GetUserPhonenumber(code string) (output *PhoneNumberOutput, err error) {
	token, err := c.AccessToken()
	if err != nil {
		return nil, err
	}

	err = c.doPost(buildAPI(getPhonePath, token), map[string]string{
		"code": code,
	}, &output)
	if err != nil {
		return
	}

	return output, output.Error()
}

func (c *Client) doRequest(url string) ([]byte, error) {
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) doPost(url string, payload, bind interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal data with error: %v", err)
	}

	resp, err := c.Post(url, "content-type/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body with error: %v", err)
	}

	return json.Unmarshal(body, &bind)
}

func buildAPI(api string, params ...interface{}) string {
	return fmt.Sprintf(endpoint+api, params...)
}
