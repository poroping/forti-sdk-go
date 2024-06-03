package models

const SystemStandaloneClusterPath = "system/standalone-cluster/"

type SystemStandaloneCluster struct {
	ClusterPeer       *[]SystemStandaloneClusterClusterPeer `json:"cluster-peer,omitempty"`
	Encryption        *string                               `json:"encryption,omitempty"`
	GroupMemberId     *int64                                `json:"group-member-id,omitempty"`
	Layer2Connection  *string                               `json:"layer2-connection,omitempty"`
	Psksecret         *string                               `json:"psksecret,omitempty"`
	SessionSyncDev    *string                               `json:"session-sync-dev,omitempty"`
	StandaloneGroupId *int64                                `json:"standalone-group-id,omitempty"`
}

type SystemStandaloneClusterClusterPeer struct {
	DownIntfsBeforeSessSync *[]SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync `json:"down-intfs-before-sess-sync,omitempty"`
	HbInterval              *int64                                                       `json:"hb-interval,omitempty"`
	HbLostThreshold         *int64                                                       `json:"hb-lost-threshold,omitempty"`
	IpsecTunnelSync         *string                                                      `json:"ipsec-tunnel-sync,omitempty"`
	Peerip                  *string                                                      `json:"peerip,omitempty"`
	Peervd                  *string                                                      `json:"peervd,omitempty"`
	SecondaryAddIpsecRoutes *string                                                      `json:"secondary-add-ipsec-routes,omitempty"`
	SessionSyncFilter       *SystemStandaloneClusterClusterPeerSessionSyncFilter         `json:"session-sync-filter,omitempty"`
	SyncId                  *int64                                                       `json:"sync-id,omitempty"`
	Syncvd                  *[]SystemStandaloneClusterClusterPeerSyncvd                  `json:"syncvd,omitempty"`
}

type SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync struct {
	Name *string `json:"name,omitempty"`
}

type SystemStandaloneClusterClusterPeerSessionSyncFilter struct {
	CustomService *[]SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService `json:"custom-service,omitempty"`
	Dstaddr       *string                                                             `json:"dstaddr,omitempty"`
	Dstaddr6      *string                                                             `json:"dstaddr6,omitempty"`
	Dstintf       *string                                                             `json:"dstintf,omitempty"`
	Srcaddr       *string                                                             `json:"srcaddr,omitempty"`
	Srcaddr6      *string                                                             `json:"srcaddr6,omitempty"`
	Srcintf       *string                                                             `json:"srcintf,omitempty"`
}

type SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService struct {
	DstPortRange *string `json:"dst-port-range,omitempty"`
	Id           *int64  `json:"id,omitempty"`
	SrcPortRange *string `json:"src-port-range,omitempty"`
}

type SystemStandaloneClusterClusterPeerSyncvd struct {
	Name *string `json:"name,omitempty"`
}
