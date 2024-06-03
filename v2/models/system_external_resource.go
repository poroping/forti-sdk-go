package models

const SystemExternalResourcePath = "system/external-resource/"

type SystemExternalResource struct {
	Category              *int64  `json:"category,omitempty"`
	Comments              *string `json:"comments,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Password              *string `json:"password,omitempty"`
	RefreshRate           *int64  `json:"refresh-rate,omitempty"`
	Resource              *string `json:"resource,omitempty"`
	ServerIdentityCheck   *string `json:"server-identity-check,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	Status                *string `json:"status,omitempty"`
	Type                  *string `json:"type,omitempty"`
	UpdateMethod          *string `json:"update-method,omitempty"`
	UserAgent             *string `json:"user-agent,omitempty"`
	Username              *string `json:"username,omitempty"`
	Uuid                  *string `json:"uuid,omitempty"`
}
