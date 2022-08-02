package models

const SystemVdomLinkPath = "system/vdom-link/"

type SystemVdomLink struct {
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
	Vcluster *string `json:"vcluster,omitempty"`
}
