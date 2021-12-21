package models

const UserLocalPath = "user/local/"

type UserLocal struct {
	AuthConcurrentOverride  *string `json:"auth-concurrent-override,omitempty"`
	AuthConcurrentValue     *int64  `json:"auth-concurrent-value,omitempty"`
	Authtimeout             *int64  `json:"authtimeout,omitempty"`
	EmailTo                 *string `json:"email-to,omitempty"`
	Fortitoken              *string `json:"fortitoken,omitempty"`
	Fosid                   *int64  `json:"fosid,omitempty"`
	LdapServer              *string `json:"ldap-server,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Passwd                  *string `json:"passwd,omitempty"`
	PasswdPolicy            *string `json:"passwd-policy,omitempty"`
	PasswdTime              *string `json:"passwd-time,omitempty"`
	PpkIdentity             *string `json:"ppk-identity,omitempty"`
	PpkSecret               *string `json:"ppk-secret,omitempty"`
	RadiusServer            *string `json:"radius-server,omitempty"`
	SmsCustomServer         *string `json:"sms-custom-server,omitempty"`
	SmsPhone                *string `json:"sms-phone,omitempty"`
	SmsServer               *string `json:"sms-server,omitempty"`
	Status                  *string `json:"status,omitempty"`
	TacacsServer            *string `json:"tacacs+-server,omitempty"`
	TwoFactor               *string `json:"two-factor,omitempty"`
	TwoFactorAuthentication *string `json:"two-factor-authentication,omitempty"`
	TwoFactorNotification   *string `json:"two-factor-notification,omitempty"`
	Type                    *string `json:"type,omitempty"`
	UsernameCaseSensitivity *string `json:"username-case-sensitivity,omitempty"`
	UsernameSensitivity     *string `json:"username-sensitivity,omitempty"`
	Workstation             *string `json:"workstation,omitempty"`
}
