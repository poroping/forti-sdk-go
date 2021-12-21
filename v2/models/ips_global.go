package models

const IpsGlobalPath = "ips/global/"

type IpsGlobal struct {
	AnomalyMode             *string                    `json:"anomaly-mode,omitempty"`
	CpAccelMode             *string                    `json:"cp-accel-mode,omitempty"`
	Database                *string                    `json:"database,omitempty"`
	DeepAppInspDbLimit      *float64                   `json:"deep-app-insp-db-limit,omitempty"`
	DeepAppInspTimeout      *float64                   `json:"deep-app-insp-timeout,omitempty"`
	EngineCount             *float64                   `json:"engine-count,omitempty"`
	ExcludeSignatures       *string                    `json:"exclude-signatures,omitempty"`
	FailOpen                *string                    `json:"fail-open,omitempty"`
	IntelligentMode         *string                    `json:"intelligent-mode,omitempty"`
	IpsReserveCpu           *string                    `json:"ips-reserve-cpu,omitempty"`
	NgfwMaxScanRange        *float64                   `json:"ngfw-max-scan-range,omitempty"`
	NpAccelMode             *string                    `json:"np-accel-mode,omitempty"`
	PacketLogQueueDepth     *float64                   `json:"packet-log-queue-depth,omitempty"`
	SessionLimitMode        *string                    `json:"session-limit-mode,omitempty"`
	SkypeClientPublicIpaddr *string                    `json:"skype-client-public-ipaddr,omitempty"`
	SocketSize              *float64                   `json:"socket-size,omitempty"`
	SyncSessionTtl          *string                    `json:"sync-session-ttl,omitempty"`
	TlsActiveProbe          *[]IpsGlobalTlsActiveProbe `json:"tls-active-probe,omitempty"`
	TrafficSubmit           *string                    `json:"traffic-submit,omitempty"`
}

type IpsGlobalTlsActiveProbe struct {
	Interface             *string `json:"interface,omitempty"`
	InterfaceSelectMethod *string `json:"interface-select-method,omitempty"`
	SourceIp              *string `json:"source-ip,omitempty"`
	SourceIp6             *string `json:"source-ip6,omitempty"`
	Vdom                  *string `json:"vdom,omitempty"`
}
