package models

const SwitchControllerinitialConfigTemplatePath = "switch-controller.initial-config/template/"

type SwitchControllerinitialConfigTemplate struct {
	Allowaccess *string  `json:"allowaccess,omitempty"`
	AutoIp      *string  `json:"auto-ip,omitempty"`
	DhcpServer  *string  `json:"dhcp-server,omitempty"`
	Ip          *string  `json:"ip,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Vlanid      *float64 `json:"vlanid,omitempty"`
}
