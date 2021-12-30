package models

const IcapServerPath = "icap/server/"

type IcapServer struct {
	IpAddress      *string `json:"ip-address,omitempty"`
	IpVersion      *string `json:"ip-version,omitempty"`
	Ip6Address     *string `json:"ip6-address,omitempty"`
	MaxConnections *int64  `json:"max-connections,omitempty"`
	Name           *string `json:"name,omitempty"`
	Port           *int64  `json:"port,omitempty"`
	Secure         *string `json:"secure,omitempty"`
	SslCert        *string `json:"ssl-cert,omitempty"`
}
