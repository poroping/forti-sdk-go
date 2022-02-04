package models

const WirelessControllerArrpProfilePath = "wireless-controller/arrp-profile/"

type WirelessControllerArrpProfile struct {
	Comment                *string                                                `json:"comment,omitempty"`
	DarrpOptimize          *int64                                                 `json:"darrp-optimize,omitempty"`
	DarrpOptimizeSchedules *[]WirelessControllerArrpProfileDarrpOptimizeSchedules `json:"darrp-optimize-schedules,omitempty"`
	IncludeDfsChannel      *string                                                `json:"include-dfs-channel,omitempty"`
	IncludeWeatherChannel  *string                                                `json:"include-weather-channel,omitempty"`
	MonitorPeriod          *int64                                                 `json:"monitor-period,omitempty"`
	Name                   *string                                                `json:"name,omitempty"`
	OverrideDarrpOptimize  *string                                                `json:"override-darrp-optimize,omitempty"`
	SelectionPeriod        *int64                                                 `json:"selection-period,omitempty"`
	ThresholdAp            *int64                                                 `json:"threshold-ap,omitempty"`
	ThresholdChannelLoad   *int64                                                 `json:"threshold-channel-load,omitempty"`
	ThresholdNoiseFloor    *string                                                `json:"threshold-noise-floor,omitempty"`
	ThresholdRxErrors      *int64                                                 `json:"threshold-rx-errors,omitempty"`
	ThresholdSpectralRssi  *string                                                `json:"threshold-spectral-rssi,omitempty"`
	ThresholdTxRetries     *int64                                                 `json:"threshold-tx-retries,omitempty"`
	WeightChannelLoad      *int64                                                 `json:"weight-channel-load,omitempty"`
	WeightDfsChannel       *int64                                                 `json:"weight-dfs-channel,omitempty"`
	WeightManagedAp        *int64                                                 `json:"weight-managed-ap,omitempty"`
	WeightNoiseFloor       *int64                                                 `json:"weight-noise-floor,omitempty"`
	WeightRogueAp          *int64                                                 `json:"weight-rogue-ap,omitempty"`
	WeightSpectralRssi     *int64                                                 `json:"weight-spectral-rssi,omitempty"`
	WeightWeatherChannel   *int64                                                 `json:"weight-weather-channel,omitempty"`
}

type WirelessControllerArrpProfileDarrpOptimizeSchedules struct {
	Name *string `json:"name,omitempty"`
}
