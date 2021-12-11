package models

const UserDomainControllerPath = "user/domain-controller/"

type UserDomainController struct {
	AdMode                *string                           `json:"ad-mode,omitempty"`
	AdldsDn               *string                           `json:"adlds-dn,omitempty"`
	AdldsIpAddress        *string                           `json:"adlds-ip-address,omitempty"`
	AdldsIp6              *string                           `json:"adlds-ip6,omitempty"`
	AdldsPort             *int64                            `json:"adlds-port,omitempty"`
	DnsSrvLookup          *string                           `json:"dns-srv-lookup,omitempty"`
	DomainName            *string                           `json:"domain-name,omitempty"`
	ExtraServer           []UserDomainControllerExtraServer `json:"extra-server,omitempty"`
	Hostname              *string                           `json:"hostname,omitempty"`
	Interface             *string                           `json:"interface,omitempty"`
	InterfaceSelectMethod *string                           `json:"interface-select-method,omitempty"`
	IpAddress             *string                           `json:"ip-address,omitempty"`
	Ip6                   *string                           `json:"ip6,omitempty"`
	LdapServer            []UserDomainControllerLdapServer  `json:"ldap-server,omitempty"`
	Name                  *string                           `json:"name,omitempty"`
	Password              *string                           `json:"password,omitempty"`
	Port                  *int64                            `json:"port,omitempty"`
	ReplicationPort       *int64                            `json:"replication-port,omitempty"`
	SourceIpAddress       *string                           `json:"source-ip-address,omitempty"`
	SourceIp6             *string                           `json:"source-ip6,omitempty"`
	SourcePort            *int64                            `json:"source-port,omitempty"`
	Username              *string                           `json:"username,omitempty"`
}

type UserDomainControllerExtraServer struct {
	Id              *int64  `json:"id,omitempty"`
	IpAddress       *string `json:"ip-address,omitempty"`
	Port            *int64  `json:"port,omitempty"`
	SourceIpAddress *string `json:"source-ip-address,omitempty"`
	SourcePort      *int64  `json:"source-port,omitempty"`
}

type UserDomainControllerLdapServer struct {
	Name *string `json:"name,omitempty"`
}
