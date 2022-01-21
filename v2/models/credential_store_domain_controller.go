package models

const CredentialStoreDomainControllerPath = "credential-store/domain-controller/"

type CredentialStoreDomainController struct {
	DomainName *string `json:"domain-name,omitempty"`
	Hostname   *string `json:"hostname,omitempty"`
	Ip         *string `json:"ip,omitempty"`
	Ip6        *string `json:"ip6,omitempty"`
	Password   *string `json:"password,omitempty"`
	Port       *int64  `json:"port,omitempty"`
	ServerName *string `json:"server-name,omitempty"`
	Username   *string `json:"username,omitempty"`
}
