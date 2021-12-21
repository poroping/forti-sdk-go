package models

const SystemExternalResourcePath = "system/external-resource/"

type SystemExternalResource struct {
	Category              *float64 `json:"category,omitempty"`
	Comments              *string  `json:"comments,omitempty"`
	Interface             *string  `json:"interface,omitempty"`
	InterfaceSelectMethod *string  `json:"interface-select-method,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	Password              *string  `json:"password,omitempty"`
	RefreshRate           *float64 `json:"refresh-rate,omitempty"`
	Resource              *string  `json:"resource,omitempty"`
	SourceIp              *string  `json:"source-ip,omitempty"`
	Status                *string  `json:"status,omitempty"`
	Type                  *string  `json:"type,omitempty"`
	UserAgent             *string  `json:"user-agent,omitempty"`
	Username              *string  `json:"username,omitempty"`
	Uuid                  *string  `json:"uuid,omitempty"`
}
