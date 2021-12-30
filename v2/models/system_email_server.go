package models

const SystemEmailServerPath = "system/email-server/"

type SystemEmailServer struct {
	Authenticate          *string `json:"authenticate,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Password              *string `json:"password,omitempty"`
	Port                  *int64  `json:"port,omitempty"`
	ReplyTo               *string `json:"reply-to,omitempty"`
	Security              *string `json:"security,omitempty"`
	Server                *string `json:"server,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	SourceIp6             *string `json:"source-ip6,omitempty"`
	SslMinProtoVersion    *string `json:"ssl-min-proto-version,omitempty"`
	Type                  *string `json:"type,omitempty"`
	Username              *string `json:"username,omitempty"`
	ValidateServer        *string `json:"validate-server,omitempty"`
}
