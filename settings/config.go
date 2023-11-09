package settings

import (
	"github.com/swanwish/go-common/config"
	"github.com/swanwish/go-common/logs"
)

var (
	ConfigFilePath = "conf/app.ini"
	AppId          = ""
	AppKey         = ""
)

func LoadAppSetting() {
	config.Load(ConfigFilePath)

	if appId, err := config.Get("app_id"); err == nil && appId != "" {
		AppId = appId
	}

	if appKey, err := config.Get("app_key"); err == nil && appKey != "" {
		AppKey = appKey
	}

	logs.Debugf("The app id is %s nad app key is %s", AppId, AppKey)

	if logLevel, err := config.Get("log_level"); err == nil {
		logs.SetLogLevel(logLevel)
	}
}
