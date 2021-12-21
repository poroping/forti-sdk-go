package models

const SystemVdomNetflowPath = "system/vdom-netflow/"

type SystemVdomNetflow struct {
	CollectorIp           *string  `json:"collector-ip,omitempty"`
	CollectorPort         *float64 `json:"collector-port,omitempty"`
	Interface             *string  `json:"interface,omitempty"`
	InterfaceSelectMethod *string  `json:"interface-select-method,omitempty"`
	SourceIp              *string  `json:"source-ip,omitempty"`
	VdomNetflow           *string  `json:"vdom-netflow,omitempty"`
}
