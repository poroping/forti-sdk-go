package models

const SwitchControllerQosIpDscpMapPath = "switch-controller.qos/ip-dscp-map/"

type SwitchControllerQosIpDscpMap struct {
	Description *string                            `json:"description,omitempty"`
	Map         *[]SwitchControllerQosIpDscpMapMap `json:"map,omitempty"`
	Name        *string                            `json:"name,omitempty"`
}

const SwitchControllerQosIpDscpMapMapPath = "switch-controller.qos/ip-dscp-map/map/"

type SwitchControllerQosIpDscpMapMap struct {
	CosQueue     *int64  `json:"cos-queue,omitempty"`
	Diffserv     *string `json:"diffserv,omitempty"`
	IpPrecedence *string `json:"ip-precedence,omitempty"`
	Name         *string `json:"name,omitempty"`
	Value        *string `json:"value,omitempty"`
}
