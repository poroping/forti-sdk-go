package models

const WirelessControllerInterControllerPath = "wireless-controller/inter-controller/"

type WirelessControllerInterController struct {
	FastFailoverMax     *float64                                                `json:"fast-failover-max,omitempty"`
	FastFailoverWait    *float64                                                `json:"fast-failover-wait,omitempty"`
	InterControllerKey  *string                                                 `json:"inter-controller-key,omitempty"`
	InterControllerMode *string                                                 `json:"inter-controller-mode,omitempty"`
	InterControllerPeer *[]WirelessControllerInterControllerInterControllerPeer `json:"inter-controller-peer,omitempty"`
	InterControllerPri  *string                                                 `json:"inter-controller-pri,omitempty"`
}

type WirelessControllerInterControllerInterControllerPeer struct {
	Id           *float64 `json:"id,omitempty"`
	PeerIp       *string  `json:"peer-ip,omitempty"`
	PeerPort     *float64 `json:"peer-port,omitempty"`
	PeerPriority *string  `json:"peer-priority,omitempty"`
}
