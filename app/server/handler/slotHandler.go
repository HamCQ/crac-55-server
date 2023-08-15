package handler

import (
	"crac55/app/server/code"
	"crac55/app/server/entities"
	"net/http"
)

// Slot slot信息
func (c *CracHandler) Slot(w http.ResponseWriter, r *http.Request) {
	var (
		res []entities.SlotStation
	)
	res, err := c.Service.Slot()
	if err != nil {
		code.SystemError(w, code.FailMsg, code.SystemErrCode)
		return
	}
	code.SuccussJSON(w, res)
}
