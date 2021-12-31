package models

const SystemDscpBasedPriorityPath = "system/dscp-based-priority/"

type SystemDscpBasedPriority struct {
	Ds       *int64  `json:"ds,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Priority *string `json:"priority,omitempty"`
}
