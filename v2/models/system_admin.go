package models

const SystemAdminPath = "system/admin/"

type SystemAdmin struct {
	Accprofile              *string                      `json:"accprofile,omitempty"`
	AccprofileOverride      *string                      `json:"accprofile-override,omitempty"`
	AllowRemoveAdminSession *string                      `json:"allow-remove-admin-session,omitempty"`
	Comments                *string                      `json:"comments,omitempty"`
	EmailTo                 *string                      `json:"email-to,omitempty"`
	ForcePasswordChange     *string                      `json:"force-password-change,omitempty"`
	Fortitoken              *string                      `json:"fortitoken,omitempty"`
	GuestAuth               *string                      `json:"guest-auth,omitempty"`
	GuestLang               *string                      `json:"guest-lang,omitempty"`
	GuestUsergroups         []SystemAdminGuestUsergroups `json:"guest-usergroups,omitempty"`
	Ip6Trusthost1           *string                      `json:"ip6-trusthost1,omitempty"`
	Ip6Trusthost10          *string                      `json:"ip6-trusthost10,omitempty"`
	Ip6Trusthost2           *string                      `json:"ip6-trusthost2,omitempty"`
	Ip6Trusthost3           *string                      `json:"ip6-trusthost3,omitempty"`
	Ip6Trusthost4           *string                      `json:"ip6-trusthost4,omitempty"`
	Ip6Trusthost5           *string                      `json:"ip6-trusthost5,omitempty"`
	Ip6Trusthost6           *string                      `json:"ip6-trusthost6,omitempty"`
	Ip6Trusthost7           *string                      `json:"ip6-trusthost7,omitempty"`
	Ip6Trusthost8           *string                      `json:"ip6-trusthost8,omitempty"`
	Ip6Trusthost9           *string                      `json:"ip6-trusthost9,omitempty"`
	Name                    *string                      `json:"name,omitempty"`
	Password                *string                      `json:"password,omitempty"`
	PasswordExpire          *string                      `json:"password-expire,omitempty"`
	PeerAuth                *string                      `json:"peer-auth,omitempty"`
	PeerGroup               *string                      `json:"peer-group,omitempty"`
	RadiusVdomOverride      *string                      `json:"radius-vdom-override,omitempty"`
	RemoteAuth              *string                      `json:"remote-auth,omitempty"`
	RemoteGroup             *string                      `json:"remote-group,omitempty"`
	Schedule                *string                      `json:"schedule,omitempty"`
	SmsCustomServer         *string                      `json:"sms-custom-server,omitempty"`
	SmsPhone                *string                      `json:"sms-phone,omitempty"`
	SmsServer               *string                      `json:"sms-server,omitempty"`
	SshCertificate          *string                      `json:"ssh-certificate,omitempty"`
	SshPublicKey1           *string                      `json:"ssh-public-key1,omitempty"`
	SshPublicKey2           *string                      `json:"ssh-public-key2,omitempty"`
	SshPublicKey3           *string                      `json:"ssh-public-key3,omitempty"`
	Trusthost1              *string                      `json:"trusthost1,omitempty"`
	Trusthost10             *string                      `json:"trusthost10,omitempty"`
	Trusthost2              *string                      `json:"trusthost2,omitempty"`
	Trusthost3              *string                      `json:"trusthost3,omitempty"`
	Trusthost4              *string                      `json:"trusthost4,omitempty"`
	Trusthost5              *string                      `json:"trusthost5,omitempty"`
	Trusthost6              *string                      `json:"trusthost6,omitempty"`
	Trusthost7              *string                      `json:"trusthost7,omitempty"`
	Trusthost8              *string                      `json:"trusthost8,omitempty"`
	Trusthost9              *string                      `json:"trusthost9,omitempty"`
	TwoFactor               *string                      `json:"two-factor,omitempty"`
	TwoFactorAuthentication *string                      `json:"two-factor-authentication,omitempty"`
	TwoFactorNotification   *string                      `json:"two-factor-notification,omitempty"`
	Vdom                    []SystemAdminVdom            `json:"vdom,omitempty"`
	Wildcard                *string                      `json:"wildcard,omitempty"`
}

type SystemAdminGuestUsergroups struct {
	Name *string `json:"name,omitempty"`
}

type SystemAdminVdom struct {
	Name *string `json:"name,omitempty"`
}
