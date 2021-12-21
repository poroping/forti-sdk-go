package models

const RouterCommunityListPath = "router/community-list/"

type RouterCommunityList struct {
	Name *string                    `json:"name,omitempty"`
	Rule *[]RouterCommunityListRule `json:"rule,omitempty"`
	Type *string                    `json:"type,omitempty"`
}

type RouterCommunityListRule struct {
	Action *string  `json:"action,omitempty"`
	Id     *float64 `json:"id,omitempty"`
	Match  *string  `json:"match,omitempty"`
	Regexp *string  `json:"regexp,omitempty"`
}
