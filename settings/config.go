package settings

import (
	"wayshare.cn/common/config"
	"wayshare.cn/common/logs"
)

var (
	ConfigFilePath = "conf/app.ini"
)

func LoadAppSetting() {
	config.Load(ConfigFilePath)

	if logLevel, err := config.Get("log_level"); err == nil {
		logs.SetLogLevel(logLevel)
	}
}
