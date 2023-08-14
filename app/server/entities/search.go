package entities

import "fmt"

// BxCra 开发便利
var BxCra = []string{
	"B0CRA", "B1CRA", "B2CRA", "B3CRA", "B4CRA",
	"B5CRA", "B6CRA", "B7CRA", "B8CRA", "B9CRA",
}

const (
	ModeCW    = "CW"
	ModePhone = "PHONE"
	ModeDigi  = "DIGI"
	BandM160  = "160M"
	BandM80   = "80M"
	BandM40   = "40M"
	BandM30   = "30M"
	BandM20   = "20M"
	BandM17   = "17M"
	BandM15   = "15M"
	BandM12   = "12M"
	BandM10   = "10M"
	BandM6    = "6M"
)

// SearchBxCraPointMap 初始化首页搜索得分数据
func SearchBxCraPointMap() []SearchBxCraInfo {
	var t []SearchBxCraInfo
	for i := 0; i <= 9; i++ {
		t = append(t, SearchBxCraInfo{
			CallsignStation: fmt.Sprintf("B%dCRA", i),
		})
	}
	return t
}

// SearchBxCraDetailMap 首页查询-qso概要
func SearchBxCraDetailMap() []SearchStationDetail {
	var t []SearchStationDetail
	for i := 0; i <= 9; i++ {
		t = append(t, SearchStationDetail{
			CallsignStation: fmt.Sprintf("B%dCRA", i),
		})
	}
	return t
}

// SearchBxCraInfo 0-9 信息
type SearchBxCraInfo struct {
	CallsignStation string            `json:"callsign_station"` //B0-9Cra
	Phone           SearchCommonPoint `json:"phone"`
	CW              SearchCommonPoint `json:"cw"`
	Digi            SearchCommonPoint `json:"digi"`
}

// SearchCommonPoint 首页查询-得分
type SearchCommonPoint struct {
	M160 int `json:"160M"` //160M
	M80  int `json:"80M"`
	M40  int `json:"40M"`
	M30  int `json:"30M"`
	M20  int `json:"20M"`
	M17  int `json:"17M"`
	M15  int `json:"15M"`
	M12  int `json:"12M"`
	M10  int `json:"10M"`
	M6   int `json:"6M"`
}

// SearchStationDetail 首页查询-站台归类
type SearchStationDetail struct {
	CallsignStation string `json:"callsign_station"`
	SearchModeDetail
}

// SearchModeDetail 首页查询-mode归类
type SearchModeDetail struct {
	Phone SearchQsoBand `json:"phone"`
	CW    SearchQsoBand `json:"cw"`
	Digi  SearchQsoBand `json:"digi"`
}

// SearchQsoBand 首页查询-band归类
type SearchQsoBand struct {
	M160 []SearchQsoDetail `json:"160M"` //160M
	M80  []SearchQsoDetail `json:"80M"`
	M40  []SearchQsoDetail `json:"40M"`
	M30  []SearchQsoDetail `json:"30M"`
	M20  []SearchQsoDetail `json:"20M"`
	M17  []SearchQsoDetail `json:"17M"`
	M15  []SearchQsoDetail `json:"15M"`
	M12  []SearchQsoDetail `json:"12M"`
	M10  []SearchQsoDetail `json:"10M"`
	M6   []SearchQsoDetail `json:"6M"`
}

// SearchQsoDetail 首页查询-qso信息概要
type SearchQsoDetail struct {
	Call     string `json:"call"`
	Oprator  string `json:"oprator"`
	Frequecy string `json:"frequecy"`
	QsoDate  string `json:"qso_date"`
}
