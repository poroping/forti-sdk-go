package models

const SystemVdomPath = "system/vdom/"

type SystemVdom struct {
	Flag       *float64 `json:"flag,omitempty"`
	Name       *string  `json:"name,omitempty"`
	ShortName  *string  `json:"short-name,omitempty"`
	VclusterId *float64 `json:"vcluster-id,omitempty"`
}
