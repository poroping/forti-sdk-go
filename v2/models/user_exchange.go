package models

const UserExchangePath = "user/exchange/"

type UserExchange struct {
	AuthLevel          *string             `json:"auth-level,omitempty"`
	AuthType           *string             `json:"auth-type,omitempty"`
	AutoDiscoverKdc    *string             `json:"auto-discover-kdc,omitempty"`
	ConnectProtocol    *string             `json:"connect-protocol,omitempty"`
	DomainName         *string             `json:"domain-name,omitempty"`
	HttpAuthType       *string             `json:"http-auth-type,omitempty"`
	Ip                 *string             `json:"ip,omitempty"`
	KdcIp              []UserExchangeKdcIp `json:"kdc-ip,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Password           *string             `json:"password,omitempty"`
	ServerName         *string             `json:"server-name,omitempty"`
	SslMinProtoVersion *string             `json:"ssl-min-proto-version,omitempty"`
	Username           *string             `json:"username,omitempty"`
}

type UserExchangeKdcIp struct {
	Ipv4 *string `json:"ipv4,omitempty"`
}
