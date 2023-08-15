package tools

import (
	"crac55/common/clog"
	"fmt"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
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

// CheckYear 简单检验年份是否正常
func CheckYear(y string) string {
	if !strings.HasPrefix(y, "2") {
		return strconv.Itoa(time.Now().Year())
	}
	if utf8.RuneCountInString(y) != 4 {
		return strconv.Itoa(time.Now().Year())
	}
	return y
}

// CheckRankPageParam 处理页码
func CheckRankPageParam(page string) int {
	p, err := strconv.Atoi(page)
	if p < 0 || err != nil {
		p = 0
	}
	return p
}

// CheckRankType 校验首页排序类型
func CheckRankType(t string) int {
	p, err := strconv.Atoi(t)
	//目前只有4种类型
	if p <= 0 || p > 4 || err != nil {
		p = 1
	}
	return p
}

// GetByCode 提取所在区号
func GetByCode(callsign string) (int, error) {
	var valid = regexp.MustCompile("[0-9]")
	areaCode := valid.FindString(callsign)
	return strconv.Atoi(areaCode)
}

// IsBnCra 校验是否属于 BnCRA
func IsBnCra(call string) bool {
	call = strings.ToUpper(call)
	match, err := regexp.MatchString("B[0-9]CRA", call)
	if err != nil {
		return true
	}
	return match
}

// CateMode 整理mode
func CateMode(mode string) string {
	mode = strings.ToLower(mode)
	if mode == "am" || mode == "fm" || mode == "ssb" {
		return "PHONE"
	}
	if mode == "cw" {
		return "CW"
	}
	if mode == "rtty" || mode == "psk31" ||
		mode == "ft8" || mode == "ft4" ||
		mode == "jt65" {
		return "DIGI"
	}
	return ""
}

// AwardTypeString
func AwardTypeString(t int8) string {
	switch t {
	case 1:
		return "Gold"
	case 2:
		return "Silver"
	case 3:
		return "Bronze"
	default:
		return "Unknown"
	}
}
