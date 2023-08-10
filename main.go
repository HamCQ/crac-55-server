package main

import (
	"crac55/app/server"
	"crac55/bootstap"
	"crac55/common/clog"
	"crac55/common/conf"
	"net/http"
)

func main() {
	bootstap.Init()
	clog.Log().Infoln("starting server at " + conf.AppConf.App.HostPort)
	if err := http.ListenAndServe(conf.AppConf.App.HostPort, server.Router()); err != nil {
		clog.Log().Errorln(err)
	}
}
