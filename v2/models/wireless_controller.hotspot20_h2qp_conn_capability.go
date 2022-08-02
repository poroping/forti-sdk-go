package models

const WirelessControllerHotspot20H2qpConnCapabilityPath = "wireless-controller.hotspot20/h2qp-conn-capability/"

type WirelessControllerHotspot20H2qpConnCapability struct {
	EspPort     *string `json:"esp-port,omitempty"`
	FtpPort     *string `json:"ftp-port,omitempty"`
	HttpPort    *string `json:"http-port,omitempty"`
	IcmpPort    *string `json:"icmp-port,omitempty"`
	Ikev2Port   *string `json:"ikev2-port,omitempty"`
	Ikev2XxPort *string `json:"ikev2-xx-port,omitempty"`
	Name        *string `json:"name,omitempty"`
	PptpVpnPort *string `json:"pptp-vpn-port,omitempty"`
	SshPort     *string `json:"ssh-port,omitempty"`
	TlsPort     *string `json:"tls-port,omitempty"`
	VoipTcpPort *string `json:"voip-tcp-port,omitempty"`
	VoipUdpPort *string `json:"voip-udp-port,omitempty"`
}
