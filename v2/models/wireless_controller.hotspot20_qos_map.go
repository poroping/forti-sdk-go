package models

const WirelessControllerhotspot20QosMapPath = "wireless-controller/hotspot20/qos-map/"

type WirelessControllerhotspot20QosMap struct {
	DscpExcept *[]WirelessControllerhotspot20QosMapDscpExcept `json:"dscp-except,omitempty"`
	DscpRange  *[]WirelessControllerhotspot20QosMapDscpRange  `json:"dscp-range,omitempty"`
	Name       *string                                        `json:"name,omitempty"`
}

type WirelessControllerhotspot20QosMapDscpExcept struct {
	Dscp  *int64 `json:"dscp,omitempty"`
	Index *int64 `json:"index,omitempty"`
	Up    *int64 `json:"up,omitempty"`
}

type WirelessControllerhotspot20QosMapDscpRange struct {
	High  *int64 `json:"high,omitempty"`
	Index *int64 `json:"index,omitempty"`
	Low   *int64 `json:"low,omitempty"`
	Up    *int64 `json:"up,omitempty"`
}
