# wechat

A WeChat mini app authoriation client.

### Installtion

Use `go get github.com/songjiayang/wechat` to install.

### Usage

```golang
package main

import (
	"fmt"
	"os"

	"github.com/songjiayang/wechat"
)

func main() {
    // wechat config.
	cfg := &wechat.Config{
		AppID:     os.Getenv("WECHAT_APPID"),
		SecretKey: os.Getenv("WECHAT_SECRETKEY"),
		Timeout:   30,
	}

	client := wechat.NewWechat(cfg)

	// mini app login with user code.
	output, _ := client.Login("code")
    fmt.Println(output)
    
    // get accessToken and cached.
	token, _ := client.AccessToken()
	fmt.Println(token)
}
```


### Reference

- [login api](https://mp.weixin.qq.com/debug/wxadoc/dev/api/api-login.html) 
- [get access_token](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183)


