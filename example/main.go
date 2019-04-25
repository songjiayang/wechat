package main

import (
	"fmt"
	"os"

	"github.com/songjiayang/wechat"
	"github.com/songjiayang/wechat/api"
)

func main() {
	cfg := &wechat.Config{
		AppID:     os.Getenv("WECHAT_APPID"),
		SecretKey: os.Getenv("WECHAT_SECRETKEY"),
		Timeout:   30,
	}

	client := wechat.NewWechat(cfg)

	// get accessToken and cached.
	token, _ := client.AccessToken()
	fmt.Println(token)

	// WeChat mini programs login code.
	output, _ := client.Login("code")
	fmt.Println(output)

	phoneOutput, _ := api.GetPhoneNumber(
		os.Getenv("WECHAT_APPID"),
		output.SessionKey,
		os.Getenv("WECHAT_IV"),
		os.Getenv("WECHAT_ENCRYPTED_DATA"),
	)

	fmt.Println(phoneOutput.PhoneNumber)
}
