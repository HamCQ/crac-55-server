package handler

import (
	"crac55/app/server/code"
	"crac55/app/server/entities"
	"crac55/common/clog"
	"net/http"
)

// AnalyseTotal 返回统计数值
func (c *CracHandler) AnalyseTotal(w http.ResponseWriter, r *http.Request) {
	var (
		res   entities.CommonResp
		total entities.AnalyseTotal
	)
	logCount, err := c.Service.Dao.CracLogCountAll(r.Context())
	if err != nil {
		clog.Log().Errorln(err)
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	total.LogNum = logCount
	res.Data = total
	code.RespJSON(w, res)
}
