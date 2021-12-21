package models

const SystemDscpBasedPriorityPath = "system/dscp-based-priority/"

type SystemDscpBasedPriority struct {
	Ds       *float64 `json:"ds,omitempty"`
	Fosid    *float64 `json:"fosid,omitempty"`
	Priority *string  `json:"priority,omitempty"`
}
