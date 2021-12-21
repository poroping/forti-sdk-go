package models

const WirelessControllerhotspot20H2qpWanMetricPath = "wireless-controller.hotspot20/h2qp-wan-metric/"

type WirelessControllerhotspot20H2qpWanMetric struct {
	DownlinkLoad            *float64 `json:"downlink-load,omitempty"`
	DownlinkSpeed           *float64 `json:"downlink-speed,omitempty"`
	LinkAtCapacity          *string  `json:"link-at-capacity,omitempty"`
	LinkStatus              *string  `json:"link-status,omitempty"`
	LoadMeasurementDuration *float64 `json:"load-measurement-duration,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	SymmetricWanLink        *string  `json:"symmetric-wan-link,omitempty"`
	UplinkLoad              *float64 `json:"uplink-load,omitempty"`
	UplinkSpeed             *float64 `json:"uplink-speed,omitempty"`
}
