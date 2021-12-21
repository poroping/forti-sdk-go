package models

const SystemTosBasedPriorityPath = "system/tos-based-priority/"

type SystemTosBasedPriority struct {
	Fosid    *float64 `json:"fosid,omitempty"`
	Priority *string  `json:"priority,omitempty"`
	Tos      *float64 `json:"tos,omitempty"`
}
