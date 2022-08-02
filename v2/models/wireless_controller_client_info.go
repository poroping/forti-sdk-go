package models

const WirelessControllerClientInfoPath = "wireless-controller/client-info/"

type WirelessControllerClientInfo struct {
	Vfid *string `json:"<vfid>,omitempty"`
}
