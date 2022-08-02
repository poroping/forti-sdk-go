package models

const FirewallDecryptedTrafficMirrorPath = "firewall/decrypted-traffic-mirror/"

type FirewallDecryptedTrafficMirror struct {
	Dstmac        *string                                    `json:"dstmac,omitempty"`
	Interface     *[]FirewallDecryptedTrafficMirrorInterface `json:"interface,omitempty"`
	Name          *string                                    `json:"name,omitempty"`
	TrafficSource *string                                    `json:"traffic-source,omitempty"`
	TrafficType   *string                                    `json:"traffic-type,omitempty"`
}

const FirewallDecryptedTrafficMirrorInterfacePath = "firewall/decrypted-traffic-mirror/interface/"

type FirewallDecryptedTrafficMirrorInterface struct {
	Name *string `json:"name,omitempty"`
}
