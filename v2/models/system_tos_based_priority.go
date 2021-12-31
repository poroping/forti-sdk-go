package models

const SystemTosBasedPriorityPath = "system/tos-based-priority/"

type SystemTosBasedPriority struct {
	Id       *int64  `json:"id,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Tos      *int64  `json:"tos,omitempty"`
}
