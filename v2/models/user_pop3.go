package models

const UserPop3Path = "user/pop3/"

type UserPop3 struct {
	Name               *string  `json:"name,omitempty"`
	Port               *float64 `json:"port,omitempty"`
	Secure             *string  `json:"secure,omitempty"`
	Server             *string  `json:"server,omitempty"`
	SslMinProtoVersion *string  `json:"ssl-min-proto-version,omitempty"`
}
