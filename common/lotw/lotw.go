package lotw

import (
	"crac55/common/clog"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	//queryLogUrl 查询接口
	queryLogUrl = "https://lotw.arrl.org/lotwuser/lotwreport.adi"
)

// QueryLogParam 查询参数
type QueryLogParam struct {
	Username  string
	Password  string
	StartDate string
	EndDate   string
}

var ErrInvalidAccount = errors.New("账户信息有误")

// QueryLog 查询lotw日志
func QueryLog(p QueryLogParam) ([]byte, error) {
	params := url.Values{}
	Url, err := url.Parse(queryLogUrl)
	if err != nil {
		clog.Log().Errorln(err)
		return nil, err
	}

	params.Set("login", p.Username)
	params.Set("password", p.Password)
	params.Set("qso_query", "1")
	params.Set("qso_qsl", "no")
	params.Set("qso_mydetail", "yes")
	params.Set("qso_startdate", p.StartDate)
	params.Set("qso_enddate", p.EndDate)

	Url.RawQuery = params.Encode()

	resp, err := http.Get(Url.String())
	if err != nil {
		clog.Log().Errorln(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		clog.Log().Errorln(err)
		return nil, err
	}

	headerTypeStatus := resp.Header.Get("Content-Type")
	if !strings.Contains(headerTypeStatus, "x-arrl-adif") {
		clog.Log().Errorln(p.Username + "账户信息有误")
		return nil, ErrInvalidAccount
	}

	return body, nil
}
