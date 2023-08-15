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

// AnalyseRankTopContent 排名统计
type AnalyseRankContent struct {
	Callsign string `json:"callsign"`
	Score    int    `json:"score"`
}
