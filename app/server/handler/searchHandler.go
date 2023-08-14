package handler

import (
	"crac55/app/server/code"
	"crac55/app/server/entities"
	"crac55/common/tools"
	"net/http"
)

// Search 首页搜索
func (c *CracHandler) Search(w http.ResponseWriter, r *http.Request) {
	var (
		callsign string
		year     string
		res      entities.CommonResp
	)
	values := r.URL.Query()
	callsign = values.Get("callsign")
	year = values.Get("year")
	if callsign == "" {
		code.SuccussJSON(w, res)
		return
	}
	year = tools.CheckYear(year)
	abstract := c.Service.Search(callsign, year)
	res.Data = abstract
	code.SuccussJSON(w, res)
}
