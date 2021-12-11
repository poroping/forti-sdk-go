package models

const UserTacacsPath = "user/tacacs+/"

type UserTacacs struct {
	AuthenType            *string `json:"authen-type,omitempty"`
	Authorization         *string `json:"authorization,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Key                   *string `json:"key,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Port                  *int64  `json:"port,omitempty"`
	SecondaryKey          *string `json:"secondary-key,omitempty"`
	SecondaryServer       *string `json:"secondary-server,omitempty"`
	Server                *string `json:"server,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	TertiaryKey           *string `json:"tertiary-key,omitempty"`
	TertiaryServer        *string `json:"tertiary-server,omitempty"`
}
