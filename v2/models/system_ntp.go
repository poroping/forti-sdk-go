package models

const SystemNtpPath = "system/ntp/"

type SystemNtp struct {
	Authentication *string               `json:"authentication,omitempty"`
	Interface      *[]SystemNtpInterface `json:"interface,omitempty"`
	Key            *string               `json:"key,omitempty"`
	KeyId          *int64                `json:"key-id,omitempty"`
	KeyType        *string               `json:"key-type,omitempty"`
	Ntpserver      *[]SystemNtpNtpserver `json:"ntpserver,omitempty"`
	Ntpsync        *string               `json:"ntpsync,omitempty"`
	ServerMode     *string               `json:"server-mode,omitempty"`
	SourceIp       *string               `json:"source-ip,omitempty"`
	SourceIp6      *string               `json:"source-ip6,omitempty"`
	Syncinterval   *int64                `json:"syncinterval,omitempty"`
	Type           *string               `json:"type,omitempty"`
}

const SystemNtpInterfacePath = "system/ntp/interface/"

type SystemNtpInterface struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}

const SystemNtpNtpserverPath = "system/ntp/ntpserver/"

type SystemNtpNtpserver struct {
	Authentication        *string `json:"authentication,omitempty"`
	Id                    *int64  `json:"id,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Key                   *string `json:"key,omitempty"`
	KeyId                 *int64  `json:"key-id,omitempty"`
	Ntpv3                 *string `json:"ntpv3,omitempty"`
	Server                *string `json:"server,omitempty"`
}
