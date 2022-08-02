package models

const SwitchControllerTrafficSnifferPath = "switch-controller/traffic-sniffer/"

type SwitchControllerTrafficSniffer struct {
	ErspanIp   *string                                     `json:"erspan-ip,omitempty"`
	Mode       *string                                     `json:"mode,omitempty"`
	TargetIp   *[]SwitchControllerTrafficSnifferTargetIp   `json:"target-ip,omitempty"`
	TargetMac  *[]SwitchControllerTrafficSnifferTargetMac  `json:"target-mac,omitempty"`
	TargetPort *[]SwitchControllerTrafficSnifferTargetPort `json:"target-port,omitempty"`
}

const SwitchControllerTrafficSnifferTargetIpPath = "switch-controller/traffic-sniffer/target-ip/"

type SwitchControllerTrafficSnifferTargetIp struct {
	Description *string `json:"description,omitempty"`
	Ip          *string `json:"ip,omitempty"`
}

const SwitchControllerTrafficSnifferTargetMacPath = "switch-controller/traffic-sniffer/target-mac/"

type SwitchControllerTrafficSnifferTargetMac struct {
	Description *string `json:"description,omitempty"`
	Mac         *string `json:"mac,omitempty"`
}

const SwitchControllerTrafficSnifferTargetPortPath = "switch-controller/traffic-sniffer/target-port/"

type SwitchControllerTrafficSnifferTargetPort struct {
	Description *string                                             `json:"description,omitempty"`
	InPorts     *[]SwitchControllerTrafficSnifferTargetPortInPorts  `json:"in-ports,omitempty"`
	OutPorts    *[]SwitchControllerTrafficSnifferTargetPortOutPorts `json:"out-ports,omitempty"`
	SwitchId    *string                                             `json:"switch-id,omitempty"`
}

const SwitchControllerTrafficSnifferTargetPortInPortsPath = "switch-controller/traffic-sniffer/target-port/in-ports/"

type SwitchControllerTrafficSnifferTargetPortInPorts struct {
	Name *string `json:"name,omitempty"`
}

const SwitchControllerTrafficSnifferTargetPortOutPortsPath = "switch-controller/traffic-sniffer/target-port/out-ports/"

type SwitchControllerTrafficSnifferTargetPortOutPorts struct {
	Name *string `json:"name,omitempty"`
}
