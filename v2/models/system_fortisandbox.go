package models

const SystemFortisandboxPath = "system/fortisandbox/"

type SystemFortisandbox struct {
	Email                 *string `json:"email,omitempty"`
	EncAlgorithm          *string `json:"enc-algorithm,omitempty"`
	Forticloud            *string `json:"forticloud,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Server                *string `json:"server,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	SslMinProtoVersion    *string `json:"ssl-min-proto-version,omitempty"`
	Status                *string `json:"status,omitempty"`
}
