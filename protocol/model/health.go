package model

/*
 * 请求/响应数据模型
 * 请求/响应模型跟协议无关，跟具体业务实现逻辑无关，仅用于请求/响应处理
 * 请求模型中form标签用于支持github.com/go-playground/form解码
 * 请求模型中resp标签用于支持github.com/wencan/copier深拷贝，不过一般直接使用字段名称即可
 *
 * wencan
 * 2019-06-24
 */

import "encoding/json"

type HealthServiceStatus int

const (
	HealthServiceStatusUnknown HealthServiceStatus = iota
	HealthServiceStatusServing
	HealthServiceStatusNotServing
	HealthServiceStatusServiceUnknown
)

var _HealthServiceStatusNames = []string{"unknown", "serving", "notServing", "serviceUnknown"}

type HealthCheckRequest struct {
	Service string `form:"service" json:"service"`
}

type HealthCheckResponse struct {
	Status HealthServiceStatus `json:"-"`
}

func (resp *HealthCheckResponse) MarshalJSON() ([]byte, error) {
	type Alias HealthCheckResponse
	response := &struct {
		*Alias
		Status string `json:"status"`
	}{
		Alias:  (*Alias)(resp),
		Status: _HealthServiceStatusNames[resp.Status],
	}

	return json.Marshal(response)
}
