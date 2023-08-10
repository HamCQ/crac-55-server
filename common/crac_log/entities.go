package craclog

// SuccessInfo 成功日志
type SuccessInfo struct {
	Status  bool
	Message string
}

// ErrInfo 错误信息
type ErrInfo struct {
	Line    int
	Message string
	Status  bool
}

// CQLog 从日志获取主要的信息
type CQLog struct {
	Call             string //被呼 呼号 *
	Mode             string //模式*
	Band             string //米波段*
	QSODate          string //QSO日期
	TimeOn           string
	TimeOff          string
	QSODateTimestamp int64  //QSO日期时间戳格式
	Frequency        string //频率
	StationCallsign  string //操作台呼号
	Operator         string //操作员*
	FileName         string //来源文件
	QslRcvd          string
	CNByCode         int8 //0-9
	CnRegionCode     int  //行政区代码，福建35 以此类推
}
