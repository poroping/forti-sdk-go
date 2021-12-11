package models

const UserSecurityExemptListPath = "user/security-exempt-list/"

type UserSecurityExemptList struct {
	Description *string                      `json:"description,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Rule        []UserSecurityExemptListRule `json:"rule,omitempty"`
}

type UserSecurityExemptListRule struct {
	Dstaddr []UserSecurityExemptListRuleDstaddr `json:"dstaddr,omitempty"`
	Id      *int64                              `json:"id,omitempty"`
	Service []UserSecurityExemptListRuleService `json:"service,omitempty"`
	Srcaddr []UserSecurityExemptListRuleSrcaddr `json:"srcaddr,omitempty"`
}

type UserSecurityExemptListRuleDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type UserSecurityExemptListRuleService struct {
	Name *string `json:"name,omitempty"`
}

type UserSecurityExemptListRuleSrcaddr struct {
	Name *string `json:"name,omitempty"`
}
