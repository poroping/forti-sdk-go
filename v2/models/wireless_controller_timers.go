package models

const WirelessControllerTimersPath = "wireless-controller/timers/"

type WirelessControllerTimers struct {
	BleScanReportIntv     *int64 `json:"ble-scan-report-intv,omitempty"`
	ClientIdleTimeout     *int64 `json:"client-idle-timeout,omitempty"`
	DiscoveryInterval     *int64 `json:"discovery-interval,omitempty"`
	DrmaInterval          *int64 `json:"drma-interval,omitempty"`
	EchoInterval          *int64 `json:"echo-interval,omitempty"`
	FakeApLog             *int64 `json:"fake-ap-log,omitempty"`
	IpsecIntfCleanup      *int64 `json:"ipsec-intf-cleanup,omitempty"`
	RadioStatsInterval    *int64 `json:"radio-stats-interval,omitempty"`
	RogueApLog            *int64 `json:"rogue-ap-log,omitempty"`
	StaCapabilityInterval *int64 `json:"sta-capability-interval,omitempty"`
	StaLocateTimer        *int64 `json:"sta-locate-timer,omitempty"`
	StaStatsInterval      *int64 `json:"sta-stats-interval,omitempty"`
	VapStatsInterval      *int64 `json:"vap-stats-interval,omitempty"`
}
