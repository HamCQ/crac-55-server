package handler

import (
	"crac55/app/server/code"
	"crac55/app/server/entities"
	"crac55/common/tools"
	"net/http"
)

// Award 奖状信息查询
func (c *CracHandler) Award(w http.ResponseWriter, r *http.Request) {
	var (
		callsign string
		year     string
		res      entities.AwardDetail
	)
	values := r.URL.Query()
	callsign = values.Get("callsign")
	year = tools.CheckYear(values.Get("year"))
	if callsign == "" {
		code.SuccussJSON(w, res)
		return
	}
	res, err := c.Service.Award(callsign, year)
	if err != nil {
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	code.SuccussJSON(w, res)
}
