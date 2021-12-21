package models

const VpnsslsettingsAuthenticationRulePath = "vpn.ssl.settings/authentication-rule/"

type VpnsslsettingsAuthenticationRule struct {
	Auth                 *string                                            `json:"auth,omitempty"`
	Cipher               *string                                            `json:"cipher,omitempty"`
	ClientCert           *string                                            `json:"client-cert,omitempty"`
	Groups               *[]VpnsslsettingsAuthenticationRuleGroups          `json:"groups,omitempty"`
	Fosid                *int64                                             `json:"fosid,omitempty"`
	Portal               *string                                            `json:"portal,omitempty"`
	Realm                *string                                            `json:"realm,omitempty"`
	SourceAddress        *[]VpnsslsettingsAuthenticationRuleSourceAddress   `json:"source-address,omitempty"`
	SourceAddressNegate  *string                                            `json:"source-address-negate,omitempty"`
	SourceAddress6       *[]VpnsslsettingsAuthenticationRuleSourceAddress6  `json:"source-address6,omitempty"`
	SourceAddress6Negate *string                                            `json:"source-address6-negate,omitempty"`
	SourceInterface      *[]VpnsslsettingsAuthenticationRuleSourceInterface `json:"source-interface,omitempty"`
	UserPeer             *string                                            `json:"user-peer,omitempty"`
	Users                *[]VpnsslsettingsAuthenticationRuleUsers           `json:"users,omitempty"`
}

type VpnsslsettingsAuthenticationRuleGroups struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslsettingsAuthenticationRuleSourceAddress struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslsettingsAuthenticationRuleSourceAddress6 struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslsettingsAuthenticationRuleSourceInterface struct {
	Name *string `json:"name,omitempty"`
}

type VpnsslsettingsAuthenticationRuleUsers struct {
	Name *string `json:"name,omitempty"`
}
