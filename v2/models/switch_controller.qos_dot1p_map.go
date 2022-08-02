package models

const SwitchControllerQosDot1pMapPath = "switch-controller.qos/dot1p-map/"

type SwitchControllerQosDot1pMap struct {
	Description      *string `json:"description,omitempty"`
	EgressPriTagging *string `json:"egress-pri-tagging,omitempty"`
	Name             *string `json:"name,omitempty"`
	Priority0        *string `json:"priority-0,omitempty"`
	Priority1        *string `json:"priority-1,omitempty"`
	Priority2        *string `json:"priority-2,omitempty"`
	Priority3        *string `json:"priority-3,omitempty"`
	Priority4        *string `json:"priority-4,omitempty"`
	Priority5        *string `json:"priority-5,omitempty"`
	Priority6        *string `json:"priority-6,omitempty"`
	Priority7        *string `json:"priority-7,omitempty"`
}
