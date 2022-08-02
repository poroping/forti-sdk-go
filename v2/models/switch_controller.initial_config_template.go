package models

const SwitchControllerInitialConfigTemplatePath = "switch-controller.initial-config/template/"

type SwitchControllerInitialConfigTemplate struct {
	Allowaccess *string `json:"allowaccess,omitempty"`
	AutoIp      *string `json:"auto-ip,omitempty"`
	DhcpServer  *string `json:"dhcp-server,omitempty"`
	Ip          *string `json:"ip,omitempty"`
	Name        *string `json:"name,omitempty"`
	Vlanid      *int64  `json:"vlanid,omitempty"`
}
