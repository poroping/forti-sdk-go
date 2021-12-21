package models

const WanoptPeerPath = "wanopt/peer/"

type WanoptPeer struct {
	Ip         *string `json:"ip,omitempty"`
	PeerHostId *string `json:"peer-host-id,omitempty"`
}
