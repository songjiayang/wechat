# wechat

A WeChat mini app authoriation client.

### Installtion

Use `go get github.com/songjiayang/wechat` to install.

### Usage

```golang
package main

import (
	"fmt"
	"github.com/songjiayang/wechat/api"
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
```


### Reference

- [login api](https://mp.weixin.qq.com/debug/wxadoc/dev/api/api-login.html) 
- [get access_token](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183)
- [api/getPhoneNumber](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html)


