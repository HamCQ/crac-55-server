package entities

// CommonResp 统一消息返回
type CommonResp struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
