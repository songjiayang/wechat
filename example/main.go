package main

import (
	"fmt"
	"os"

	"github.com/songjiayang/wechat"
)

func main() {
	cfg := &wechat.Config{
		AppID:     os.Getenv("WECHAT_APPID"),
		SecretKey: os.Getenv("WECHAT_SECRETKEY"),
		Timeout:   30,
	}

	client := wechat.NewClient(cfg)

	// get accessToken and cached.
	token, _ := client.AccessToken()
	fmt.Printf("token: %s \n", token)

	// WeChat mini programs login code.
	output, err := client.Login(os.Getenv("WECHAT_CODE"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("openid: %s \n", output.OpenID)

	info, err := client.GetUserPhonenumber(os.Getenv("WECHAT_PHONE_CODE"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("phone: %s \n", info.PhoneNumber())
}
