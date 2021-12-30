package models

const SwitchControllerLldpProfilePath = "switch-controller/lldp-profile/"

type SwitchControllerLldpProfile struct {
	The8021Tlvs           *string                                          `json:"802.1-tlvs,omitempty"`
	The8023Tlvs           *string                                          `json:"802.3-tlvs,omitempty"`
	AutoIsl               *string                                          `json:"auto-isl,omitempty"`
	AutoIslHelloTimer     *int64                                           `json:"auto-isl-hello-timer,omitempty"`
	AutoIslPortGroup      *int64                                           `json:"auto-isl-port-group,omitempty"`
	AutoIslReceiveTimeout *int64                                           `json:"auto-isl-receive-timeout,omitempty"`
	AutoMclagIcl          *string                                          `json:"auto-mclag-icl,omitempty"`
	CustomTlvs            *[]SwitchControllerLldpProfileCustomTlvs         `json:"custom-tlvs,omitempty"`
	MedLocationService    *[]SwitchControllerLldpProfileMedLocationService `json:"med-location-service,omitempty"`
	MedNetworkPolicy      *[]SwitchControllerLldpProfileMedNetworkPolicy   `json:"med-network-policy,omitempty"`
	MedTlvs               *string                                          `json:"med-tlvs,omitempty"`
	Name                  *string                                          `json:"name,omitempty"`
}

type SwitchControllerLldpProfileCustomTlvs struct {
	InformationString *string `json:"information-string,omitempty"`
	Name              *string `json:"name,omitempty"`
	Oui               *string `json:"oui,omitempty"`
	Subtype           *int64  `json:"subtype,omitempty"`
}

type SwitchControllerLldpProfileMedLocationService struct {
	Name          *string `json:"name,omitempty"`
	Status        *string `json:"status,omitempty"`
	SysLocationId *string `json:"sys-location-id,omitempty"`
}

type SwitchControllerLldpProfileMedNetworkPolicy struct {
	AssignVlan *string `json:"assign-vlan,omitempty"`
	Dscp       *int64  `json:"dscp,omitempty"`
	Name       *string `json:"name,omitempty"`
	Priority   *int64  `json:"priority,omitempty"`
	Status     *string `json:"status,omitempty"`
	VlanIntf   *string `json:"vlan-intf,omitempty"`
}
