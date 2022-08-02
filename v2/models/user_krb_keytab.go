package models

const UserKrbKeytabPath = "user/krb-keytab/"

type UserKrbKeytab struct {
	Keytab     *string                    `json:"keytab,omitempty"`
	LdapServer *[]UserKrbKeytabLdapServer `json:"ldap-server,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	PacData    *string                    `json:"pac-data,omitempty"`
	Principal  *string                    `json:"principal,omitempty"`
}

const UserKrbKeytabLdapServerPath = "user/krb-keytab/ldap-server/"

type UserKrbKeytabLdapServer struct {
	Name *string `json:"name,omitempty"`
}
