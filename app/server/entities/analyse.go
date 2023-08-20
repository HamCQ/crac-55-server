package entities

const (
	CnDiffCra     = 1
	CnCra         = 2
	GlobleDiffCra = 3
	GlobleCra     = 4
)

// AnalyseTotal 首页总参数
type AnalyseTotal struct {
	LogNum        int64 `json:"log_num"`
	AwardNum      int64 `json:"award_num"`
	SingleCallNum int64 `json:"single_call_num"`
}

type AnalyseRankRes struct {
	Total     int64                `json:"total"`
	Page      int                  `json:"page"`
	PageSize  int                  `json:"page_size"`
	PageCount int                  `json:"page_count"`
	Data      []AnalyseRankContent `json:"data"`
}

// AnalyseRankTopContent 排名统计
type AnalyseRankContent struct {
	Number   int    `json:"number"`
	Callsign string `json:"callsign"`
	Score    int    `json:"score"`
}

// AnalyseBy09 BY电台通联数量统计(0-9区)
type AnalyseBy09 struct {
	ByCode   int   `json:"by_code"`
	CWNum    int64 `json:"cw_num"`
	PhoneNum int64 `json:"phone_num"`
	DigiNum  int64 `json:"digi_num"`
}

// AnalyseProvince 省份QSO数量统计
type AnalyseProvince struct {
	CnRegionCode int   `json:"cn_region_code"` //行政区域编码
	QsoNum       int64 `json:"qso_num"`
}

type AnalyseBnCra struct {
	CallsignStation string `json:"callsign_station"`
	CWNum           int64  `json:"cw_num"`
	PhoneNum        int64  `json:"phone_num"`
	DigiNum         int64  `json:"digi_num"`
	Sum             int64  `json:"sum"`
}
