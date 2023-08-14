package handler

import (
	"crac55/app/server/code"
	"crac55/common/tools"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// AnalyseTotal 返回统计数值
func (c *CracHandler) AnalyseTotal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, ok := vars["year"]
	if !ok {
		year = strconv.Itoa(time.Now().Year())
	}
	year = tools.CheckYear(year)
	num, err := c.Service.AnalyseNum(year)
	if err != nil {
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	code.SuccussJSON(w, num)
}
