package models

const AuthenticationSchemePath = "authentication/scheme/"

type AuthenticationScheme struct {
	DomainController *string                             `json:"domain-controller,omitempty"`
	FssoAgentForNtlm *string                             `json:"fsso-agent-for-ntlm,omitempty"`
	FssoGuest        *string                             `json:"fsso-guest,omitempty"`
	KerberosKeytab   *string                             `json:"kerberos-keytab,omitempty"`
	Method           *string                             `json:"method,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	NegotiateNtlm    *string                             `json:"negotiate-ntlm,omitempty"`
	RequireTfa       *string                             `json:"require-tfa,omitempty"`
	SamlServer       *string                             `json:"saml-server,omitempty"`
	SamlTimeout      *int64                              `json:"saml-timeout,omitempty"`
	SshCa            *string                             `json:"ssh-ca,omitempty"`
	UserCert         *string                             `json:"user-cert,omitempty"`
	UserDatabase     *[]AuthenticationSchemeUserDatabase `json:"user-database,omitempty"`
}

const AuthenticationSchemeUserDatabasePath = "authentication/scheme/user-database/"

type AuthenticationSchemeUserDatabase struct {
	Name *string `json:"name,omitempty"`
}
