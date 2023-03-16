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
	output, err := client.Login(os.Getenv("WECHAT_CODE"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

	if os.Getenv("WECHAT_IV") != "" {
		phoneOutput, _ := api.GetPhoneNumber(
			os.Getenv("WECHAT_APPID"),
			output.SessionKey,
			os.Getenv("WECHAT_IV"),
			os.Getenv("WECHAT_ENCRYPTED_DATA"),
		)
		fmt.Println(phoneOutput.PhoneNumber)
	}
}
