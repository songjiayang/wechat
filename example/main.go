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

	client := wechat.NewWechat(cfg)

	// get accessToken and cached.
	token, _ := client.AccessToken()
	fmt.Println(token)

	// WeChat mini app login code.
	output, _ := client.Login("code")
	fmt.Println(output)
}
