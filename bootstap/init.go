package bootstap

import (
	"crac55/app/server/code"
	"crac55/common/cache"
	"crac55/common/clog"
	"crac55/common/conf"
	"crac55/common/db"
)

// Init 初始化资源
func Init() {
	conf.Init()
	clog.InitLog()
	db.Init()
	cache.Init()
	code.InitI18n()
}
