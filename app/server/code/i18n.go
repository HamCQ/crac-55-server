package code

import (
	"crac55/common/clog"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// InitI18n 初始化i18n
func InitI18n() {
	bundle = i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFile("app/server/code/i18n/zh-CN.zh.toml")
	if err != nil {
		clog.Log().Errorln("i8n Chinese文件加载失败", err)
		return
	}
	_, err = bundle.LoadMessageFile("app/server/code/i18n/en-US.es.toml")
	if err != nil {
		clog.Log().Errorln("i18n English文件加载失败", err)
		return
	}
}
