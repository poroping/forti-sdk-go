package models

const SwitchControllerAutoConfigPolicyPath = "switch-controller.auto-config/policy/"

type SwitchControllerAutoConfigPolicy struct {
	IgmpFloodReport    *string `json:"igmp-flood-report,omitempty"`
	IgmpFloodTraffic   *string `json:"igmp-flood-traffic,omitempty"`
	Name               *string `json:"name,omitempty"`
	PoeStatus          *string `json:"poe-status,omitempty"`
	QosPolicy          *string `json:"qos-policy,omitempty"`
	StormControlPolicy *string `json:"storm-control-policy,omitempty"`
}
