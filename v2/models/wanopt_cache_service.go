package models

const WanoptCacheServicePath = "wanopt/cache-service/"

type WanoptCacheService struct {
	AcceptableConnections *string                      `json:"acceptable-connections,omitempty"`
	Collaboration         *string                      `json:"collaboration,omitempty"`
	DeviceId              *string                      `json:"device-id,omitempty"`
	DstPeer               *[]WanoptCacheServiceDstPeer `json:"dst-peer,omitempty"`
	PreferScenario        *string                      `json:"prefer-scenario,omitempty"`
	SrcPeer               *[]WanoptCacheServiceSrcPeer `json:"src-peer,omitempty"`
}

type WanoptCacheServiceDstPeer struct {
	AuthType   *int64  `json:"auth-type,omitempty"`
	DeviceId   *string `json:"device-id,omitempty"`
	EncodeType *int64  `json:"encode-type,omitempty"`
	Ip         *string `json:"ip,omitempty"`
	Priority   *int64  `json:"priority,omitempty"`
}

type WanoptCacheServiceSrcPeer struct {
	AuthType   *int64  `json:"auth-type,omitempty"`
	DeviceId   *string `json:"device-id,omitempty"`
	EncodeType *int64  `json:"encode-type,omitempty"`
	Ip         *string `json:"ip,omitempty"`
	Priority   *int64  `json:"priority,omitempty"`
}
