package model

/*
 * 请求/响应数据模型
 * 请求/响应模型跟协议无关，跟具体业务实现逻辑无关，仅用于请求/响应处理
 * 请求模型中form标签用于支持github.com/go-playground/form解码
 * 请求模型中validate标签用于支持github.com/go-playground/validator做请求参数检查
 * 请求模型中reply标签用于支持github.com/wencan/copier深拷贝
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

var (
	_HealthServiceStatus_Names = map[HealthServiceStatus]string{
		HealthServiceStatusUnknown:        "unknown",
		HealthServiceStatusServing:        "serving",
		HealthServiceStatusNotServing:     "notServing",
		HealthServiceStatusServiceUnknown: "serviceUnknown",
	}

	_HealthServiceStatus_Values = make(map[string]HealthServiceStatus)
)

func init() {
	for value, name := range _HealthServiceStatus_Names {
		_HealthServiceStatus_Values[name] = value
	}
}

func HealthServiceStatusName(status HealthServiceStatus) string {
	return _HealthServiceStatus_Names[status]
}

func HealthServiceStatusFromName(name string) HealthServiceStatus {
	status, exists := _HealthServiceStatus_Values[name]
	if !exists {
		return HealthServiceStatusUnknown
	}
	return status
}

type HealthCheckRequest struct {
	Service string `form:"service" json:"service" validate:"required"`
}

type HealthCheckResponse struct {
	Status HealthServiceStatus `json:"-" reply:"status"`
}

func (resp *HealthCheckResponse) MarshalJSON() ([]byte, error) {
	type Alias HealthCheckResponse
	response := &struct {
		*Alias
		Status string `json:"status"`
	}{
		Alias:  (*Alias)(resp),
		Status: HealthServiceStatusName(resp.Status),
	}

	return json.Marshal(response)
}

func (resp *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	type Alias HealthCheckResponse
	response := &struct {
		*Alias
		Status string `json:"status"`
	}{
		Alias: (*Alias)(resp),
	}

	err := json.Unmarshal(data, response)
	if err != nil {
		return err
	}

	response.Alias.Status = HealthServiceStatusFromName(response.Status)

	return nil
}
