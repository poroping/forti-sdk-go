package models

const UserSecurityExemptListPath = "user/security-exempt-list/"

type UserSecurityExemptList struct {
	Description *string                       `json:"description,omitempty"`
	Name        *string                       `json:"name,omitempty"`
	Rule        *[]UserSecurityExemptListRule `json:"rule,omitempty"`
}

const UserSecurityExemptListRulePath = "user/security-exempt-list/rule/"

type UserSecurityExemptListRule struct {
	Dstaddr *[]UserSecurityExemptListRuleDstaddr `json:"dstaddr,omitempty"`
	Id      *int64                               `json:"id,omitempty"`
	Service *[]UserSecurityExemptListRuleService `json:"service,omitempty"`
	Srcaddr *[]UserSecurityExemptListRuleSrcaddr `json:"srcaddr,omitempty"`
}

const UserSecurityExemptListRuleDstaddrPath = "user/security-exempt-list/rule/dstaddr/"

type UserSecurityExemptListRuleDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const UserSecurityExemptListRuleServicePath = "user/security-exempt-list/rule/service/"

type UserSecurityExemptListRuleService struct {
	Name *string `json:"name,omitempty"`
}

const UserSecurityExemptListRuleSrcaddrPath = "user/security-exempt-list/rule/srcaddr/"

type UserSecurityExemptListRuleSrcaddr struct {
	Name *string `json:"name,omitempty"`
}
