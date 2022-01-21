package models

const SystemStandaloneClusterPath = "system/standalone-cluster/"

type SystemStandaloneCluster struct {
	Encryption        *string `json:"encryption,omitempty"`
	GroupMemberId     *int64  `json:"group-member-id,omitempty"`
	Layer2Connection  *string `json:"layer2-connection,omitempty"`
	Psksecret         *string `json:"psksecret,omitempty"`
	SessionSyncDev    *string `json:"session-sync-dev,omitempty"`
	StandaloneGroupId *int64  `json:"standalone-group-id,omitempty"`
}
