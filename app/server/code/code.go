package code

import (
	"crac55/app/server/entities"
	"encoding/json"
	"io"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	_codes = map[int]struct{}{}
	bundle = i18n.NewBundle(language.Chinese)
)

// RespJSON 响应json
func RespJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	content, err := json.Marshal(data)
	if err != nil {
		SystemError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(content))
}

// SystemError 返回系统错误信息
func SystemError(w http.ResponseWriter, errMsg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	content, _ := json.Marshal(&entities.CommonResp{
		Code:    code,
		Message: errMsg,
	})
	io.WriteString(w, string(content))
}
