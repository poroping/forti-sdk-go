package models

const SwitchControllerIgmpSnoopingPath = "switch-controller/igmp-snooping/"

type SwitchControllerIgmpSnooping struct {
	AgingTime             *int64  `json:"aging-time,omitempty"`
	FloodUnknownMulticast *string `json:"flood-unknown-multicast,omitempty"`
}
