package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CacheRedis cache.CacheConf
	KvConf     kv.KvConf
	Mysql      struct {
		DataSource string
	}
	LogConf        LogConf
	GameApiConfig  GameApiConfig
	SendCodeConfig SendCodeConfig
}
type LogConf struct {
	ServiceName         string `json:",optional"`                                    // 服务名称
	Mode                string `json:",default=console,options=console|file|volume"` // 日志模式，console-输出到console，file-输出到当前服务器（容器）文件，，volume-输出docker挂在文件内
	Path                string `json:",default=logs"`                                // 日志存储路径
	Level               string `json:",default=info,options=info|error|severe"`      // 日志级别
	Compress            bool   `json:",optional"`                                    // 是否开启gzip压缩
	KeepDays            int    `json:",optional"`                                    // 日志保留天数
	StackCooldownMillis int    `json:",default=100"`                                 // 日志write间隔
}

type GameApiConfig struct {
	GameUrl   string `json:","`
	SecretId  string `json:","`
	SecretKey string `json:","`
	AppId     int    `json:","`
	SvrRegion string `json:","`
}

type SendCodeConfig struct {
	GameUrl   string `json:","`
	ApiSecret string `json:","`
}
