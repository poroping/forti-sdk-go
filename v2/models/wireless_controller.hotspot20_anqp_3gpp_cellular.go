package models

const WirelessControllerHotspot20Anqp3gppCellularPath = "wireless-controller.hotspot20/anqp-3gpp-cellular/"

type WirelessControllerHotspot20Anqp3gppCellular struct {
	MccMncList *[]WirelessControllerHotspot20Anqp3gppCellularMccMncList `json:"mcc-mnc-list,omitempty"`
	Name       *string                                                  `json:"name,omitempty"`
}

type WirelessControllerHotspot20Anqp3gppCellularMccMncList struct {
	Id  *int64  `json:"id,omitempty"`
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
}
