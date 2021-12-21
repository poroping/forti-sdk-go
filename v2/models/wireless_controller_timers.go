package models

const WirelessControllerTimersPath = "wireless-controller/timers/"

type WirelessControllerTimers struct {
	BleScanReportIntv     *float64 `json:"ble-scan-report-intv,omitempty"`
	ClientIdleTimeout     *float64 `json:"client-idle-timeout,omitempty"`
	DiscoveryInterval     *float64 `json:"discovery-interval,omitempty"`
	DrmaInterval          *float64 `json:"drma-interval,omitempty"`
	EchoInterval          *float64 `json:"echo-interval,omitempty"`
	FakeApLog             *float64 `json:"fake-ap-log,omitempty"`
	IpsecIntfCleanup      *float64 `json:"ipsec-intf-cleanup,omitempty"`
	RadioStatsInterval    *float64 `json:"radio-stats-interval,omitempty"`
	RogueApLog            *float64 `json:"rogue-ap-log,omitempty"`
	StaCapabilityInterval *float64 `json:"sta-capability-interval,omitempty"`
	StaLocateTimer        *float64 `json:"sta-locate-timer,omitempty"`
	StaStatsInterval      *float64 `json:"sta-stats-interval,omitempty"`
	VapStatsInterval      *float64 `json:"vap-stats-interval,omitempty"`
}
