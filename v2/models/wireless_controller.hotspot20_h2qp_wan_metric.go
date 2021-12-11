package models

const WirelessControllerhotspot20H2qpWanMetricPath = "wireless-controller.hotspot20/h2qp-wan-metric/"

type WirelessControllerhotspot20H2qpWanMetric struct {
	DownlinkLoad            *int64  `json:"downlink-load,omitempty"`
	DownlinkSpeed           *int64  `json:"downlink-speed,omitempty"`
	LinkAtCapacity          *string `json:"link-at-capacity,omitempty"`
	LinkStatus              *string `json:"link-status,omitempty"`
	LoadMeasurementDuration *int64  `json:"load-measurement-duration,omitempty"`
	Name                    *string `json:"name,omitempty"`
	SymmetricWanLink        *string `json:"symmetric-wan-link,omitempty"`
	UplinkLoad              *int64  `json:"uplink-load,omitempty"`
	UplinkSpeed             *int64  `json:"uplink-speed,omitempty"`
}
