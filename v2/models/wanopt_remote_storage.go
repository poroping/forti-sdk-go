package models

const WanoptRemoteStoragePath = "wanopt/remote-storage/"

type WanoptRemoteStorage struct {
	LocalCacheId  *string `json:"local-cache-id,omitempty"`
	RemoteCacheId *string `json:"remote-cache-id,omitempty"`
	RemoteCacheIp *string `json:"remote-cache-ip,omitempty"`
	Status        *string `json:"status,omitempty"`
}
