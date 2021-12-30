package models

const UserRadiusPath = "user/radius/"

type UserRadius struct {
	AccountingServer                       *[]UserRadiusAccountingServer `json:"accounting-server,omitempty"`
	AcctAllServers                         *string                       `json:"acct-all-servers,omitempty"`
	AcctInterimInterval                    *int64                        `json:"acct-interim-interval,omitempty"`
	AllUsergroup                           *string                       `json:"all-usergroup,omitempty"`
	AuthType                               *string                       `json:"auth-type,omitempty"`
	Class                                  *[]UserRadiusClass            `json:"class,omitempty"`
	GroupOverrideAttrType                  *string                       `json:"group-override-attr-type,omitempty"`
	H3cCompatibility                       *string                       `json:"h3c-compatibility,omitempty"`
	Interface                              *string                       `json:"interface,omitempty"`
	InterfaceSelectMethod                  *string                       `json:"interface-select-method,omitempty"`
	Name                                   *string                       `json:"name,omitempty"`
	NasIp                                  *string                       `json:"nas-ip,omitempty"`
	PasswordEncoding                       *string                       `json:"password-encoding,omitempty"`
	PasswordRenewal                        *string                       `json:"password-renewal,omitempty"`
	RadiusCoa                              *string                       `json:"radius-coa,omitempty"`
	RadiusPort                             *int64                        `json:"radius-port,omitempty"`
	Rsso                                   *string                       `json:"rsso,omitempty"`
	RssoContextTimeout                     *int64                        `json:"rsso-context-timeout,omitempty"`
	RssoEndpointAttribute                  *string                       `json:"rsso-endpoint-attribute,omitempty"`
	RssoEndpointBlockAttribute             *string                       `json:"rsso-endpoint-block-attribute,omitempty"`
	RssoEpOneIpOnly                        *string                       `json:"rsso-ep-one-ip-only,omitempty"`
	RssoFlushIpSession                     *string                       `json:"rsso-flush-ip-session,omitempty"`
	RssoLogFlags                           *string                       `json:"rsso-log-flags,omitempty"`
	RssoLogPeriod                          *int64                        `json:"rsso-log-period,omitempty"`
	RssoRadiusResponse                     *string                       `json:"rsso-radius-response,omitempty"`
	RssoRadiusServerPort                   *int64                        `json:"rsso-radius-server-port,omitempty"`
	RssoSecret                             *string                       `json:"rsso-secret,omitempty"`
	RssoValidateRequestSecret              *string                       `json:"rsso-validate-request-secret,omitempty"`
	SecondarySecret                        *string                       `json:"secondary-secret,omitempty"`
	SecondaryServer                        *string                       `json:"secondary-server,omitempty"`
	Secret                                 *string                       `json:"secret,omitempty"`
	Server                                 *string                       `json:"server,omitempty"`
	SourceIp                               *string                       `json:"source-ip,omitempty"`
	SsoAttribute                           *string                       `json:"sso-attribute,omitempty"`
	SsoAttributeKey                        *string                       `json:"sso-attribute-key,omitempty"`
	SsoAttributeValueOverride              *string                       `json:"sso-attribute-value-override,omitempty"`
	SwitchControllerAcctFastFramedipDetect *int64                        `json:"switch-controller-acct-fast-framedip-detect,omitempty"`
	SwitchControllerServiceType            *string                       `json:"switch-controller-service-type,omitempty"`
	TertiarySecret                         *string                       `json:"tertiary-secret,omitempty"`
	TertiaryServer                         *string                       `json:"tertiary-server,omitempty"`
	Timeout                                *int64                        `json:"timeout,omitempty"`
	UseManagementVdom                      *string                       `json:"use-management-vdom,omitempty"`
	UsernameCaseSensitive                  *string                       `json:"username-case-sensitive,omitempty"`
}

type UserRadiusAccountingServer struct {
	Id                    *int64  `json:"id,omitempty"`
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	Port                  *int64  `json:"port,omitempty"`
	Secret                *string `json:"secret,omitempty"`
	Server                *string `json:"server,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	Status                *string `json:"status,omitempty"`
}

type UserRadiusClass struct {
	Name *string `json:"name,omitempty"`
}
