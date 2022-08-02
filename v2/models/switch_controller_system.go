package models

const SwitchControllerSystemPath = "switch-controller/system/"

type SwitchControllerSystem struct {
	DataSyncInterval        *int64  `json:"data-sync-interval,omitempty"`
	DynamicPeriodicInterval *int64  `json:"dynamic-periodic-interval,omitempty"`
	IotHoldoff              *int64  `json:"iot-holdoff,omitempty"`
	IotMacIdle              *int64  `json:"iot-mac-idle,omitempty"`
	IotScanInterval         *int64  `json:"iot-scan-interval,omitempty"`
	IotWeightThreshold      *int64  `json:"iot-weight-threshold,omitempty"`
	NacPeriodicInterval     *int64  `json:"nac-periodic-interval,omitempty"`
	ParallelProcess         *int64  `json:"parallel-process,omitempty"`
	ParallelProcessOverride *string `json:"parallel-process-override,omitempty"`
	TunnelMode              *string `json:"tunnel-mode,omitempty"`
}
