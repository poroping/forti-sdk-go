package models

const UserLdapPath = "user/ldap/"

type UserLdap struct {
	AccountKeyFilter        *string  `json:"account-key-filter,omitempty"`
	AccountKeyProcessing    *string  `json:"account-key-processing,omitempty"`
	Antiphish               *string  `json:"antiphish,omitempty"`
	CaCert                  *string  `json:"ca-cert,omitempty"`
	Cnid                    *string  `json:"cnid,omitempty"`
	Dn                      *string  `json:"dn,omitempty"`
	GroupFilter             *string  `json:"group-filter,omitempty"`
	GroupMemberCheck        *string  `json:"group-member-check,omitempty"`
	GroupObjectFilter       *string  `json:"group-object-filter,omitempty"`
	GroupSearchBase         *string  `json:"group-search-base,omitempty"`
	Interface               *string  `json:"interface,omitempty"`
	InterfaceSelectMethod   *string  `json:"interface-select-method,omitempty"`
	MemberAttr              *string  `json:"member-attr,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	ObtainUserInfo          *string  `json:"obtain-user-info,omitempty"`
	Password                *string  `json:"password,omitempty"`
	PasswordAttr            *string  `json:"password-attr,omitempty"`
	PasswordExpiryWarning   *string  `json:"password-expiry-warning,omitempty"`
	PasswordRenewal         *string  `json:"password-renewal,omitempty"`
	Port                    *float64 `json:"port,omitempty"`
	SearchType              *string  `json:"search-type,omitempty"`
	SecondaryServer         *string  `json:"secondary-server,omitempty"`
	Secure                  *string  `json:"secure,omitempty"`
	Server                  *string  `json:"server,omitempty"`
	ServerIdentityCheck     *string  `json:"server-identity-check,omitempty"`
	SourceIp                *string  `json:"source-ip,omitempty"`
	SourcePort              *float64 `json:"source-port,omitempty"`
	SslMinProtoVersion      *string  `json:"ssl-min-proto-version,omitempty"`
	TertiaryServer          *string  `json:"tertiary-server,omitempty"`
	TwoFactor               *string  `json:"two-factor,omitempty"`
	TwoFactorAuthentication *string  `json:"two-factor-authentication,omitempty"`
	TwoFactorNotification   *string  `json:"two-factor-notification,omitempty"`
	Type                    *string  `json:"type,omitempty"`
	UserInfoExchangeServer  *string  `json:"user-info-exchange-server,omitempty"`
	Username                *string  `json:"username,omitempty"`
}
