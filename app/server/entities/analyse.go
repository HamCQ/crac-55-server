package entities

// AnalyseTotal 首页总参数
type AnalyseTotal struct {
	LogNum        int64 `json:"log_num"`
	AwardNum      int64 `json:"award_num"`
	SingleCallNum int64 `json:"single_call_num"`
}
