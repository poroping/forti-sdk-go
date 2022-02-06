package models

const SystemClusterSyncPath = "system/cluster-sync/"

type SystemClusterSync struct {
	DownIntfsBeforeSessSync *[]SystemClusterSyncDownIntfsBeforeSessSync `json:"down-intfs-before-sess-sync,omitempty"`
	HbInterval              *int64                                      `json:"hb-interval,omitempty"`
	HbLostThreshold         *int64                                      `json:"hb-lost-threshold,omitempty"`
	IkeHeartbeatInterval    *int64                                      `json:"ike-heartbeat-interval,omitempty"`
	IkeMonitor              *string                                     `json:"ike-monitor,omitempty"`
	IkeMonitorInterval      *int64                                      `json:"ike-monitor-interval,omitempty"`
	IkeSeqjumpSpeed         *int64                                      `json:"ike-seqjump-speed,omitempty"`
	IpsecTunnelSync         *string                                     `json:"ipsec-tunnel-sync,omitempty"`
	Peerip                  *string                                     `json:"peerip,omitempty"`
	Peervd                  *string                                     `json:"peervd,omitempty"`
	SecondaryAddIpsecRoutes *string                                     `json:"secondary-add-ipsec-routes,omitempty"`
	SessionSyncFilter       *SystemClusterSyncSessionSyncFilter         `json:"session-sync-filter,omitempty"`
	SlaveAddIkeRoutes       *string                                     `json:"slave-add-ike-routes,omitempty"`
	SyncId                  *int64                                      `json:"sync-id,omitempty"`
	Syncvd                  *[]SystemClusterSyncSyncvd                  `json:"syncvd,omitempty"`
}

type SystemClusterSyncDownIntfsBeforeSessSync struct {
	Name *string `json:"name,omitempty"`
}

type SystemClusterSyncSessionSyncFilter struct {
	CustomService *[]SystemClusterSyncSessionSyncFilterCustomService `json:"custom-service,omitempty"`
	Dstaddr       *string                                            `json:"dstaddr,omitempty"`
	Dstaddr6      *string                                            `json:"dstaddr6,omitempty"`
	Dstintf       *string                                            `json:"dstintf,omitempty"`
	Srcaddr       *string                                            `json:"srcaddr,omitempty"`
	Srcaddr6      *string                                            `json:"srcaddr6,omitempty"`
	Srcintf       *string                                            `json:"srcintf,omitempty"`
}

type SystemClusterSyncSessionSyncFilterCustomService struct {
	DstPortRange *string `json:"dst-port-range,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	SrcPortRange *string `json:"src-port-range,omitempty"`
}

type SystemClusterSyncSyncvd struct {
	Name *string `json:"name,omitempty"`
}
