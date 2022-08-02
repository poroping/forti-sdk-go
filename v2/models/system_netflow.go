package models

const SystemNetflowPath = "system/netflow/"

type SystemNetflow struct {
	ActiveFlowTimeout     *int64  `json:"active-flow-timeout,omitempty"`
	CollectorIp           *string `json:"collector-ip,omitempty"`
	CollectorPort         *int64  `json:"collector-port,omitempty"`
	InactiveFlowTimeout   *int64  `json:"inactive-flow-timeout,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	TemplateTxCounter     *int64  `json:"template-tx-counter,omitempty"`
	TemplateTxTimeout     *int64  `json:"template-tx-timeout,omitempty"`
}
