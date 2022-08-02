package models

const SwitchControllerFortilinkSettingsPath = "switch-controller/fortilink-settings/"

type SwitchControllerFortilinkSettings struct {
	Fortilink     *string                                    `json:"fortilink,omitempty"`
	InactiveTimer *int64                                     `json:"inactive-timer,omitempty"`
	LinkDownFlush *string                                    `json:"link-down-flush,omitempty"`
	NacPorts      *SwitchControllerFortilinkSettingsNacPorts `json:"nac-ports,omitempty"`
	Name          *string                                    `json:"name,omitempty"`
}

const SwitchControllerFortilinkSettingsNacPortsPath = "switch-controller/fortilink-settings/nac-ports/"

type SwitchControllerFortilinkSettingsNacPorts struct {
	LanSegment      *string                                                     `json:"lan-segment,omitempty"`
	MemberChange    *int64                                                      `json:"member-change,omitempty"`
	NacLanInterface *string                                                     `json:"nac-lan-interface,omitempty"`
	NacSegmentVlans *[]SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans `json:"nac-segment-vlans,omitempty"`
	OnboardingVlan  *string                                                     `json:"onboarding-vlan,omitempty"`
	ParentKey       *string                                                     `json:"parent-key,omitempty"`
}

const SwitchControllerFortilinkSettingsNacPortsNacSegmentVlansPath = "switch-controller/fortilink-settings/nac-ports/nac-segment-vlans/"

type SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans struct {
	VlanName *string `json:"vlan-name,omitempty"`
}
