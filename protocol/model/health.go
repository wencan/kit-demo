package model

import "encoding/json"

type HealthServiceStatus int

const (
	HealthServiceStatusUnknown HealthServiceStatus = iota
	HealthServiceStatusServing
	HealthServiceStatusNotServing
	HealthServiceStatusServiceUnknown
)

var _HealthServiceStatusNames = []string{"Unknown", "Serving", "NotServing", "ServiceUnknown"}

type HealthCheckRequest struct {
	Service string `json:"service"`
}

type HealthCheckResponse struct {
	Status HealthServiceStatus
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
