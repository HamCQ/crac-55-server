package tools

import (
	"crac55/common/clog"
	"fmt"
	"runtime/debug"
)

// Go 集中管理协程
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				clog.Log().Errorln(err)
				clog.Log().Errorln(fmt.Sprint(string(debug.Stack())))
			}
		}()
		x()
	}()
}
