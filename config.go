package wechat

import "time"

type Config struct {
	AppID     string `json:"appid"`
	SecretKey string `json:"secret_key"`
	Timeout   int    `json:"timeout"`
}

func (cfg *Config) TimeoutDuration() time.Duration {
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 30
	}

	return time.Duration(timeout) * time.Second
}
