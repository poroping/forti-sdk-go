package models

const CifsDomainControllerPath = "cifs/domain-controller/"

type CifsDomainController struct {
	DomainName *string  `json:"domain-name,omitempty"`
	Ip         *string  `json:"ip,omitempty"`
	Ip6        *string  `json:"ip6,omitempty"`
	Password   *string  `json:"password,omitempty"`
	Port       *float64 `json:"port,omitempty"`
	ServerName *string  `json:"server-name,omitempty"`
	Username   *string  `json:"username,omitempty"`
}
