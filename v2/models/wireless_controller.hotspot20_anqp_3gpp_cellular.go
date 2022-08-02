package models

const WirelessControllerHotspot20Anqp3GppCellularPath = "wireless-controller.hotspot20/anqp-3gpp-cellular/"

type WirelessControllerHotspot20Anqp3GppCellular struct {
	MccMncList *[]WirelessControllerHotspot20Anqp3GppCellularMccMncList `json:"mcc-mnc-list,omitempty"`
	Name       *string                                                  `json:"name,omitempty"`
}

const WirelessControllerHotspot20Anqp3GppCellularMccMncListPath = "wireless-controller.hotspot20/anqp-3gpp-cellular/mcc-mnc-list/"

type WirelessControllerHotspot20Anqp3GppCellularMccMncList struct {
	Id  *int64  `json:"id,omitempty"`
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
}
