package handler

import (
	"crac55/app/server/code"
	"crac55/common/tools"
	"net/http"
)

// AnalyseTotal 返回统计数值
func (c *CracHandler) AnalyseTotal(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	year := tools.CheckYear(values.Get("year"))
	num, err := c.Service.AnalyseNum(year)
	if err != nil {
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	code.SuccussJSON(w, num)
}

// AnalyseRank 排名统计-top5
func (c *CracHandler) AnalyseRankTop5(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	year := tools.CheckYear(values.Get("year"))
	res, err := c.Service.AnalyseRankTop(year)
	if err != nil {
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	code.SuccussJSON(w, res)
}

// AnalyseRankAll 排名统计-所有
func (c *CracHandler) AnalyseRankAll(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	year := tools.CheckYear(values.Get("year"))
	page := tools.CheckRankPageParam(values.Get("page"))
	rankType := tools.CheckRankType(values.Get("type"))
	res, err := c.Service.AnalyseRankAll(year, rankType, page)
	if err != nil {
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	code.SuccussJSON(w, res)
}
