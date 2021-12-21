package models

const WirelessControllerhotspot20Anqp3gppCellularPath = "wireless-controller.hotspot20/anqp-3gpp-cellular/"

type WirelessControllerhotspot20Anqp3gppCellular struct {
	MccMncList *[]WirelessControllerhotspot20Anqp3gppCellularMccMncList `json:"mcc-mnc-list,omitempty"`
	Name       *string                                                  `json:"name,omitempty"`
}

type WirelessControllerhotspot20Anqp3gppCellularMccMncList struct {
	Id  *float64 `json:"id,omitempty"`
	Mcc *string  `json:"mcc,omitempty"`
	Mnc *string  `json:"mnc,omitempty"`
}
