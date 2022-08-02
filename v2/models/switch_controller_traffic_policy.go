package models

const SwitchControllerTrafficPolicyPath = "switch-controller/traffic-policy/"

type SwitchControllerTrafficPolicy struct {
	CosQueue            *int64  `json:"cos-queue,omitempty"`
	Description         *string `json:"description,omitempty"`
	GuaranteedBandwidth *int64  `json:"guaranteed-bandwidth,omitempty"`
	GuaranteedBurst     *int64  `json:"guaranteed-burst,omitempty"`
	MaximumBurst        *int64  `json:"maximum-burst,omitempty"`
	Name                *string `json:"name,omitempty"`
	PolicerStatus       *string `json:"policer-status,omitempty"`
	Type                *string `json:"type,omitempty"`
}
