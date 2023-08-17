package bootstap

import (
	"crac55/app/cron"
	"crac55/app/server/code"
	"crac55/common/cache"
	"crac55/common/clog"
	"crac55/common/conf"
	"crac55/common/db"
	"crac55/common/limiter"
	"crac55/common/tools"
)

// Init 初始化资源
func Init() {
	conf.Init()                                       //初始化配置
	clog.InitLog()                                    //初始化日志
	db.Init()                                         //初始化数据库
	cache.Init()                                      //初始化缓存
	code.InitI18n()                                   //初始化预设状态码
	tools.Go(func() { cron.NewCrontab().InitTask() }) //初始化定时任务
	tools.Go(func() { limiter.CleanupVisitors() })
}
