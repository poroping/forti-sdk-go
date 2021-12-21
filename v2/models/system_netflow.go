package models

const SystemNetflowPath = "system/netflow/"

type SystemNetflow struct {
	ActiveFlowTimeout     *float64 `json:"active-flow-timeout,omitempty"`
	CollectorIp           *string  `json:"collector-ip,omitempty"`
	CollectorPort         *float64 `json:"collector-port,omitempty"`
	InactiveFlowTimeout   *float64 `json:"inactive-flow-timeout,omitempty"`
	Interface             *string  `json:"interface,omitempty"`
	InterfaceSelectMethod *string  `json:"interface-select-method,omitempty"`
	SourceIp              *string  `json:"source-ip,omitempty"`
	TemplateTxCounter     *float64 `json:"template-tx-counter,omitempty"`
	TemplateTxTimeout     *float64 `json:"template-tx-timeout,omitempty"`
}
