package models

const WirelessControllerhotspot20QosMapPath = "wireless-controller.hotspot20/qos-map/"

type WirelessControllerhotspot20QosMap struct {
	DscpExcept *[]WirelessControllerhotspot20QosMapDscpExcept `json:"dscp-except,omitempty"`
	DscpRange  *[]WirelessControllerhotspot20QosMapDscpRange  `json:"dscp-range,omitempty"`
	Name       *string                                        `json:"name,omitempty"`
}

type WirelessControllerhotspot20QosMapDscpExcept struct {
	Dscp  *float64 `json:"dscp,omitempty"`
	Index *float64 `json:"index,omitempty"`
	Up    *float64 `json:"up,omitempty"`
}

type WirelessControllerhotspot20QosMapDscpRange struct {
	High  *float64 `json:"high,omitempty"`
	Index *float64 `json:"index,omitempty"`
	Low   *float64 `json:"low,omitempty"`
	Up    *float64 `json:"up,omitempty"`
}
