package models

const UserPeerPath = "user/peer/"

type UserPeer struct {
	Ca                 *string `json:"ca,omitempty"`
	Cn                 *string `json:"cn,omitempty"`
	CnType             *string `json:"cn-type,omitempty"`
	LdapMode           *string `json:"ldap-mode,omitempty"`
	LdapPassword       *string `json:"ldap-password,omitempty"`
	LdapServer         *string `json:"ldap-server,omitempty"`
	LdapUsername       *string `json:"ldap-username,omitempty"`
	MandatoryCaVerify  *string `json:"mandatory-ca-verify,omitempty"`
	Name               *string `json:"name,omitempty"`
	OcspOverrideServer *string `json:"ocsp-override-server,omitempty"`
	Passwd             *string `json:"passwd,omitempty"`
	Subject            *string `json:"subject,omitempty"`
	TwoFactor          *string `json:"two-factor,omitempty"`
}
