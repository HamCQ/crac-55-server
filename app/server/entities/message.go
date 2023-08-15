package entities

// CommonResp 统一消息返回
type CommonResp struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SearchPayload 首页搜索返回
type SearchPayload struct {
	Bxcra       []SearchStationDetail `json:"bxcra"`
	Bncra       SearchStationDetail   `json:"bncra"`
	AwardStatus AwardInfo             `json:"award_info"`
}

// AnalyseRankTopRes 排名统计
type AnalyseRankTopRes struct {
	UpdateTime    string               `json:"update_time"`
	CnDiffCra     []AnalyseRankContent `json:"cn_diff_cra"`
	CnCra         []AnalyseRankContent `json:"cn_cra"`
	GlobleDiffCra []AnalyseRankContent `json:"globle_diff_cra"`
	GlobleCra     []AnalyseRankContent `json:"globle_cra"`
}
