package models

const WirelessControllerGlobalPath = "wireless-controller/global/"

type WirelessControllerGlobal struct {
	AcdProcessCount       *int64  `json:"acd-process-count,omitempty"`
	ApLogServer           *string `json:"ap-log-server,omitempty"`
	ApLogServerIp         *string `json:"ap-log-server-ip,omitempty"`
	ApLogServerPort       *int64  `json:"ap-log-server-port,omitempty"`
	ControlMessageOffload *string `json:"control-message-offload,omitempty"`
	DataEthernetII        *string `json:"data-ethernet-II,omitempty"`
	DiscoveryMcAddr       *string `json:"discovery-mc-addr,omitempty"`
	FiappEthType          *int64  `json:"fiapp-eth-type,omitempty"`
	ImageDownload         *string `json:"image-download,omitempty"`
	IpsecBaseIp           *string `json:"ipsec-base-ip,omitempty"`
	LinkAggregation       *string `json:"link-aggregation,omitempty"`
	Location              *string `json:"location,omitempty"`
	MaxClients            *int64  `json:"max-clients,omitempty"`
	MaxRetransmit         *int64  `json:"max-retransmit,omitempty"`
	MeshEthType           *int64  `json:"mesh-eth-type,omitempty"`
	NacInterval           *int64  `json:"nac-interval,omitempty"`
	Name                  *string `json:"name,omitempty"`
	RogueScanMacAdjacency *int64  `json:"rogue-scan-mac-adjacency,omitempty"`
	TunnelMode            *string `json:"tunnel-mode,omitempty"`
	WtpShare              *string `json:"wtp-share,omitempty"`
}
