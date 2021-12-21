package models

const SystemHaPath = "system/ha/"

type SystemHa struct {
	Arps                          *float64                     `json:"arps,omitempty"`
	ArpsInterval                  *float64                     `json:"arps-interval,omitempty"`
	Authentication                *string                      `json:"authentication,omitempty"`
	CpuThreshold                  *string                      `json:"cpu-threshold,omitempty"`
	Encryption                    *string                      `json:"encryption,omitempty"`
	FailoverHoldTime              *float64                     `json:"failover-hold-time,omitempty"`
	FtpProxyThreshold             *string                      `json:"ftp-proxy-threshold,omitempty"`
	GratuitousArps                *string                      `json:"gratuitous-arps,omitempty"`
	GroupId                       *float64                     `json:"group-id,omitempty"`
	GroupName                     *string                      `json:"group-name,omitempty"`
	HaDirect                      *string                      `json:"ha-direct,omitempty"`
	HaEthType                     *string                      `json:"ha-eth-type,omitempty"`
	HaMgmtInterfaces              *[]SystemHaHaMgmtInterfaces  `json:"ha-mgmt-interfaces,omitempty"`
	HaMgmtStatus                  *string                      `json:"ha-mgmt-status,omitempty"`
	HaUptimeDiffMargin            *float64                     `json:"ha-uptime-diff-margin,omitempty"`
	HbInterval                    *float64                     `json:"hb-interval,omitempty"`
	HbIntervalInMilliseconds      *string                      `json:"hb-interval-in-milliseconds,omitempty"`
	HbLostThreshold               *float64                     `json:"hb-lost-threshold,omitempty"`
	Hbdev                         *string                      `json:"hbdev,omitempty"`
	HcEthType                     *string                      `json:"hc-eth-type,omitempty"`
	HelloHolddown                 *float64                     `json:"hello-holddown,omitempty"`
	HttpProxyThreshold            *string                      `json:"http-proxy-threshold,omitempty"`
	ImapProxyThreshold            *string                      `json:"imap-proxy-threshold,omitempty"`
	InterClusterSessionSync       *string                      `json:"inter-cluster-session-sync,omitempty"`
	Key                           *string                      `json:"key,omitempty"`
	L2epEthType                   *string                      `json:"l2ep-eth-type,omitempty"`
	LinkFailedSignal              *string                      `json:"link-failed-signal,omitempty"`
	LoadBalanceAll                *string                      `json:"load-balance-all,omitempty"`
	LogicalSn                     *string                      `json:"logical-sn,omitempty"`
	MemoryBasedFailover           *string                      `json:"memory-based-failover,omitempty"`
	MemoryCompatibleMode          *string                      `json:"memory-compatible-mode,omitempty"`
	MemoryFailoverFlipTimeout     *float64                     `json:"memory-failover-flip-timeout,omitempty"`
	MemoryFailoverMonitorPeriod   *float64                     `json:"memory-failover-monitor-period,omitempty"`
	MemoryFailoverSampleRate      *float64                     `json:"memory-failover-sample-rate,omitempty"`
	MemoryFailoverThreshold       *float64                     `json:"memory-failover-threshold,omitempty"`
	MemoryThreshold               *string                      `json:"memory-threshold,omitempty"`
	Mode                          *string                      `json:"mode,omitempty"`
	Monitor                       *string                      `json:"monitor,omitempty"`
	MulticastTtl                  *float64                     `json:"multicast-ttl,omitempty"`
	NntpProxyThreshold            *string                      `json:"nntp-proxy-threshold,omitempty"`
	Override                      *string                      `json:"override,omitempty"`
	OverrideWaitTime              *float64                     `json:"override-wait-time,omitempty"`
	Password                      *string                      `json:"password,omitempty"`
	PingserverFailoverThreshold   *float64                     `json:"pingserver-failover-threshold,omitempty"`
	PingserverFlipTimeout         *float64                     `json:"pingserver-flip-timeout,omitempty"`
	PingserverMonitorInterface    *string                      `json:"pingserver-monitor-interface,omitempty"`
	PingserverSecondaryForceReset *string                      `json:"pingserver-secondary-force-reset,omitempty"`
	PingserverSlaveForceReset     *string                      `json:"pingserver-slave-force-reset,omitempty"`
	Pop3ProxyThreshold            *string                      `json:"pop3-proxy-threshold,omitempty"`
	Priority                      *float64                     `json:"priority,omitempty"`
	RouteHold                     *float64                     `json:"route-hold,omitempty"`
	RouteTtl                      *float64                     `json:"route-ttl,omitempty"`
	RouteWait                     *float64                     `json:"route-wait,omitempty"`
	Schedule                      *string                      `json:"schedule,omitempty"`
	SecondaryVcluster             *[]SystemHaSecondaryVcluster `json:"secondary-vcluster,omitempty"`
	SessionPickup                 *string                      `json:"session-pickup,omitempty"`
	SessionPickupConnectionless   *string                      `json:"session-pickup-connectionless,omitempty"`
	SessionPickupDelay            *string                      `json:"session-pickup-delay,omitempty"`
	SessionPickupExpectation      *string                      `json:"session-pickup-expectation,omitempty"`
	SessionPickupNat              *string                      `json:"session-pickup-nat,omitempty"`
	SessionSyncDev                *string                      `json:"session-sync-dev,omitempty"`
	SmtpProxyThreshold            *string                      `json:"smtp-proxy-threshold,omitempty"`
	SsdFailover                   *string                      `json:"ssd-failover,omitempty"`
	StandaloneConfigSync          *string                      `json:"standalone-config-sync,omitempty"`
	StandaloneMgmtVdom            *string                      `json:"standalone-mgmt-vdom,omitempty"`
	SyncConfig                    *string                      `json:"sync-config,omitempty"`
	SyncPacketBalance             *string                      `json:"sync-packet-balance,omitempty"`
	UnicastGateway                *string                      `json:"unicast-gateway,omitempty"`
	UnicastHb                     *string                      `json:"unicast-hb,omitempty"`
	UnicastHbNetmask              *string                      `json:"unicast-hb-netmask,omitempty"`
	UnicastHbPeerip               *string                      `json:"unicast-hb-peerip,omitempty"`
	UnicastPeers                  *[]SystemHaUnicastPeers      `json:"unicast-peers,omitempty"`
	UnicastStatus                 *string                      `json:"unicast-status,omitempty"`
	UninterruptiblePrimaryWait    *float64                     `json:"uninterruptible-primary-wait,omitempty"`
	UninterruptibleUpgrade        *string                      `json:"uninterruptible-upgrade,omitempty"`
	VclusterId                    *float64                     `json:"vcluster-id,omitempty"`
	Vcluster2                     *string                      `json:"vcluster2,omitempty"`
	Vdom                          *string                      `json:"vdom,omitempty"`
	Weight                        *string                      `json:"weight,omitempty"`
}

type SystemHaHaMgmtInterfaces struct {
	Dst       *string  `json:"dst,omitempty"`
	Gateway   *string  `json:"gateway,omitempty"`
	Gateway6  *string  `json:"gateway6,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Interface *string  `json:"interface,omitempty"`
}

type SystemHaSecondaryVcluster struct {
	Monitor                       *string  `json:"monitor,omitempty"`
	Override                      *string  `json:"override,omitempty"`
	OverrideWaitTime              *float64 `json:"override-wait-time,omitempty"`
	PingserverFailoverThreshold   *float64 `json:"pingserver-failover-threshold,omitempty"`
	PingserverMonitorInterface    *string  `json:"pingserver-monitor-interface,omitempty"`
	PingserverSecondaryForceReset *string  `json:"pingserver-secondary-force-reset,omitempty"`
	PingserverSlaveForceReset     *string  `json:"pingserver-slave-force-reset,omitempty"`
	Priority                      *float64 `json:"priority,omitempty"`
	VclusterId                    *float64 `json:"vcluster-id,omitempty"`
	Vdom                          *string  `json:"vdom,omitempty"`
}

type SystemHaUnicastPeers struct {
	Id     *float64 `json:"id,omitempty"`
	PeerIp *string  `json:"peer-ip,omitempty"`
}
