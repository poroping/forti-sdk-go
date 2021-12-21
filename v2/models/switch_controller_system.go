package models

const SwitchControllerSystemPath = "switch-controller/system/"

type SwitchControllerSystem struct {
	DataSyncInterval        *float64 `json:"data-sync-interval,omitempty"`
	DynamicPeriodicInterval *float64 `json:"dynamic-periodic-interval,omitempty"`
	IotHoldoff              *float64 `json:"iot-holdoff,omitempty"`
	IotMacIdle              *float64 `json:"iot-mac-idle,omitempty"`
	IotScanInterval         *float64 `json:"iot-scan-interval,omitempty"`
	IotWeightThreshold      *float64 `json:"iot-weight-threshold,omitempty"`
	NacPeriodicInterval     *float64 `json:"nac-periodic-interval,omitempty"`
	ParallelProcess         *float64 `json:"parallel-process,omitempty"`
	ParallelProcessOverride *string  `json:"parallel-process-override,omitempty"`
	TunnelMode              *string  `json:"tunnel-mode,omitempty"`
}
