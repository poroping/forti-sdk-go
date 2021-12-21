package models

const SystemVdomSflowPath = "system/vdom-sflow/"

type SystemVdomSflow struct {
	CollectorIp           *string  `json:"collector-ip,omitempty"`
	CollectorPort         *float64 `json:"collector-port,omitempty"`
	Interface             *string  `json:"interface,omitempty"`
	InterfaceSelectMethod *string  `json:"interface-select-method,omitempty"`
	SourceIp              *string  `json:"source-ip,omitempty"`
	VdomSflow             *string  `json:"vdom-sflow,omitempty"`
}
