package models

const UserFssoPath = "user/fsso/"

type UserFsso struct {
	GroupPollInterval     *int64  `json:"group-poll-interval,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	LdapPoll              *string `json:"ldap-poll,omitempty"`
	LdapPollFilter        *string `json:"ldap-poll-filter,omitempty"`
	LdapPollInterval      *int64  `json:"ldap-poll-interval,omitempty"`
	LdapServer            *string `json:"ldap-server,omitempty"`
	LogonTimeout          *int64  `json:"logon-timeout,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Password              *string `json:"password,omitempty"`
	Password2             *string `json:"password2,omitempty"`
	Password3             *string `json:"password3,omitempty"`
	Password4             *string `json:"password4,omitempty"`
	Password5             *string `json:"password5,omitempty"`
	Port                  *int64  `json:"port,omitempty"`
	Port2                 *int64  `json:"port2,omitempty"`
	Port3                 *int64  `json:"port3,omitempty"`
	Port4                 *int64  `json:"port4,omitempty"`
	Port5                 *int64  `json:"port5,omitempty"`
	Server                *string `json:"server,omitempty"`
	Server2               *string `json:"server2,omitempty"`
	Server3               *string `json:"server3,omitempty"`
	Server4               *string `json:"server4,omitempty"`
	Server5               *string `json:"server5,omitempty"`
	Sni                   *string `json:"sni,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	SourceIp6             *string `json:"source-ip6,omitempty"`
	Ssl                   *string `json:"ssl,omitempty"`
	SslServerHostIpCheck  *string `json:"ssl-server-host-ip-check,omitempty"`
	SslTrustedCert        *string `json:"ssl-trusted-cert,omitempty"`
	Type                  *string `json:"type,omitempty"`
	UserInfoServer        *string `json:"user-info-server,omitempty"`
}
