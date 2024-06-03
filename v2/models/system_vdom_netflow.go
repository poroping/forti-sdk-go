package models

const SystemVdomNetflowPath = "system/vdom-netflow/"

type SystemVdomNetflow struct {
	CollectorIp           *string                        `json:"collector-ip,omitempty"`
	CollectorPort         *int64                         `json:"collector-port,omitempty"`
	Collectors            *[]SystemVdomNetflowCollectors `json:"collectors,omitempty"`
	Interface             *string                        `json:"interface,omitempty"`
	InterfaceSelectMethod *string                        `json:"interface-select-method,omitempty"`
	SourceIp              *string                        `json:"source-ip,omitempty"`
	VdomNetflow           *string                        `json:"vdom-netflow,omitempty"`
}

type SystemVdomNetflowCollectors struct {
	CollectorIp           *string `json:"collector-ip,omitempty"`
	CollectorPort         *int64  `json:"collector-port,omitempty"`
	Id                    *int64  `json:"id,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
}
