package models

const SystemDscpBasedPriorityPath = "system/dscp-based-priority/"

type SystemDscpBasedPriority struct {
	Ds       *int64  `json:"ds,omitempty"`
	Fosid    *int64  `json:"fosid,omitempty"`
	Priority *string `json:"priority,omitempty"`
}
