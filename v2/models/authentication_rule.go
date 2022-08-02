package models

const AuthenticationRulePath = "authentication/rule/"

type AuthenticationRule struct {
	ActiveAuthMethod *string                       `json:"active-auth-method,omitempty"`
	Comments         *string                       `json:"comments,omitempty"`
	Dstaddr          *[]AuthenticationRuleDstaddr  `json:"dstaddr,omitempty"`
	Dstaddr6         *[]AuthenticationRuleDstaddr6 `json:"dstaddr6,omitempty"`
	IpBased          *string                       `json:"ip-based,omitempty"`
	Name             *string                       `json:"name,omitempty"`
	Protocol         *string                       `json:"protocol,omitempty"`
	Srcaddr          *[]AuthenticationRuleSrcaddr  `json:"srcaddr,omitempty"`
	Srcaddr6         *[]AuthenticationRuleSrcaddr6 `json:"srcaddr6,omitempty"`
	Srcintf          *[]AuthenticationRuleSrcintf  `json:"srcintf,omitempty"`
	SsoAuthMethod    *string                       `json:"sso-auth-method,omitempty"`
	Status           *string                       `json:"status,omitempty"`
	TransactionBased *string                       `json:"transaction-based,omitempty"`
	WebAuthCookie    *string                       `json:"web-auth-cookie,omitempty"`
	WebPortal        *string                       `json:"web-portal,omitempty"`
}

const AuthenticationRuleDstaddrPath = "authentication/rule/dstaddr/"

type AuthenticationRuleDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const AuthenticationRuleDstaddr6Path = "authentication/rule/dstaddr6/"

type AuthenticationRuleDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const AuthenticationRuleSrcaddrPath = "authentication/rule/srcaddr/"

type AuthenticationRuleSrcaddr struct {
	Name *string `json:"name,omitempty"`
}

const AuthenticationRuleSrcaddr6Path = "authentication/rule/srcaddr6/"

type AuthenticationRuleSrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

const AuthenticationRuleSrcintfPath = "authentication/rule/srcintf/"

type AuthenticationRuleSrcintf struct {
	Name *string `json:"name,omitempty"`
}
