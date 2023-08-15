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

// SearchBnCraMap 初始化首页搜索得分数据 bncra
func SearchBnCraMap() SearchStationDetail {
	return SearchStationDetail{
		CallsignStation: "BnCRA",
	}
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

// SearchStationDetail 首页查询-站台归类
type SearchStationDetail struct {
	CallsignStation string `json:"callsign_station"`
	Score           int    `json:"score"`
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
	Oprator  string `json:"oprator"`
	Frequecy string `json:"frequecy"`
}

// AwardInfo 获奖信息
type AwardInfo struct {
	Status      bool   `json:"status"`
	AwardType   int8   `json:"award_type"`
	AwardString string `json:"award_string"`
	Continent   string `json:"continent"`
}

// AwardDetail 奖状详细信息
type AwardDetail struct {
	AwardType      int8   `json:"award_type"`
	AwardString    string `json:"award_string"`
	Continent      string `json:"continent"`
	BncraNum       int64  `json:"bncra_num"`
	Callsign       string `json:"callsign"`
	ImageUrl       string `json:"img_url"`
	ImageUrlOrigin string `json:"img_url_origin"`
	Combination    int64  `json:"combination"`
}
