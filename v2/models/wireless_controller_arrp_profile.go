package models

const WirelessControllerArrpProfilePath = "wireless-controller/arrp-profile/"

type WirelessControllerArrpProfile struct {
	Comment               *string  `json:"comment,omitempty"`
	IncludeDfsChannel     *string  `json:"include-dfs-channel,omitempty"`
	IncludeWeatherChannel *string  `json:"include-weather-channel,omitempty"`
	MonitorPeriod         *float64 `json:"monitor-period,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	SelectionPeriod       *float64 `json:"selection-period,omitempty"`
	ThresholdAp           *float64 `json:"threshold-ap,omitempty"`
	ThresholdChannelLoad  *float64 `json:"threshold-channel-load,omitempty"`
	ThresholdNoiseFloor   *string  `json:"threshold-noise-floor,omitempty"`
	ThresholdRxErrors     *float64 `json:"threshold-rx-errors,omitempty"`
	ThresholdSpectralRssi *string  `json:"threshold-spectral-rssi,omitempty"`
	ThresholdTxRetries    *float64 `json:"threshold-tx-retries,omitempty"`
	WeightChannelLoad     *float64 `json:"weight-channel-load,omitempty"`
	WeightDfsChannel      *float64 `json:"weight-dfs-channel,omitempty"`
	WeightManagedAp       *float64 `json:"weight-managed-ap,omitempty"`
	WeightNoiseFloor      *float64 `json:"weight-noise-floor,omitempty"`
	WeightRogueAp         *float64 `json:"weight-rogue-ap,omitempty"`
	WeightSpectralRssi    *float64 `json:"weight-spectral-rssi,omitempty"`
	WeightWeatherChannel  *float64 `json:"weight-weather-channel,omitempty"`
}
