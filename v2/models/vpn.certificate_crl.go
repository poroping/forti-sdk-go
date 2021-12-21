package models

const VpncertificateCrlPath = "vpn.certificate/crl/"

type VpncertificateCrl struct {
	Crl            *string  `json:"crl,omitempty"`
	HttpUrl        *string  `json:"http-url,omitempty"`
	LdapPassword   *string  `json:"ldap-password,omitempty"`
	LdapServer     *string  `json:"ldap-server,omitempty"`
	LdapUsername   *string  `json:"ldap-username,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Range          *string  `json:"range,omitempty"`
	ScepCert       *string  `json:"scep-cert,omitempty"`
	ScepUrl        *string  `json:"scep-url,omitempty"`
	Source         *string  `json:"source,omitempty"`
	SourceIp       *string  `json:"source-ip,omitempty"`
	UpdateInterval *float64 `json:"update-interval,omitempty"`
	UpdateVdom     *string  `json:"update-vdom,omitempty"`
}
