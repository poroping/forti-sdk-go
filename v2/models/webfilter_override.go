package models

const WebfilterOverridePath = "webfilter/override/"

type WebfilterOverride struct {
	Expires    *string `json:"expires,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Initiator  *string `json:"initiator,omitempty"`
	Ip         *string `json:"ip,omitempty"`
	Ip6        *string `json:"ip6,omitempty"`
	NewProfile *string `json:"new-profile,omitempty"`
	OldProfile *string `json:"old-profile,omitempty"`
	Scope      *string `json:"scope,omitempty"`
	Status     *string `json:"status,omitempty"`
	User       *string `json:"user,omitempty"`
	UserGroup  *string `json:"user-group,omitempty"`
}
