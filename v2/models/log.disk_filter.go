package models

const LogdiskFilterPath = "log.disk/filter/"

type LogdiskFilter struct {
	Admin            *string                   `json:"admin,omitempty"`
	Anomaly          *string                   `json:"anomaly,omitempty"`
	Auth             *string                   `json:"auth,omitempty"`
	CpuMemoryUsage   *string                   `json:"cpu-memory-usage,omitempty"`
	Dhcp             *string                   `json:"dhcp,omitempty"`
	DlpArchive       *string                   `json:"dlp-archive,omitempty"`
	Event            *string                   `json:"event,omitempty"`
	Filter           *string                   `json:"filter,omitempty"`
	FilterType       *string                   `json:"filter-type,omitempty"`
	ForwardTraffic   *string                   `json:"forward-traffic,omitempty"`
	FreeStyle        *[]LogdiskFilterFreeStyle `json:"free-style,omitempty"`
	Gtp              *string                   `json:"gtp,omitempty"`
	Ha               *string                   `json:"ha,omitempty"`
	Ipsec            *string                   `json:"ipsec,omitempty"`
	LdbMonitor       *string                   `json:"ldb-monitor,omitempty"`
	LocalTraffic     *string                   `json:"local-traffic,omitempty"`
	MulticastTraffic *string                   `json:"multicast-traffic,omitempty"`
	Pattern          *string                   `json:"pattern,omitempty"`
	Ppp              *string                   `json:"ppp,omitempty"`
	Radius           *string                   `json:"radius,omitempty"`
	Severity         *string                   `json:"severity,omitempty"`
	SnifferTraffic   *string                   `json:"sniffer-traffic,omitempty"`
	SslvpnLogAdm     *string                   `json:"sslvpn-log-adm,omitempty"`
	SslvpnLogAuth    *string                   `json:"sslvpn-log-auth,omitempty"`
	SslvpnLogSession *string                   `json:"sslvpn-log-session,omitempty"`
	System           *string                   `json:"system,omitempty"`
	VipSsl           *string                   `json:"vip-ssl,omitempty"`
	Voip             *string                   `json:"voip,omitempty"`
	WanOpt           *string                   `json:"wan-opt,omitempty"`
	WirelessActivity *string                   `json:"wireless-activity,omitempty"`
}

type LogdiskFilterFreeStyle struct {
	Category   *string `json:"category,omitempty"`
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filter-type,omitempty"`
	Id         *int64  `json:"id,omitempty"`
}
