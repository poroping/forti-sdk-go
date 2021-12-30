package models

const SystemTosBasedPriorityPath = "system/tos-based-priority/"

type SystemTosBasedPriority struct {
	Fosid    *int64  `json:"fosid,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Tos      *int64  `json:"tos,omitempty"`
}
