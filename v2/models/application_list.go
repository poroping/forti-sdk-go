package models

const ApplicationListPath = "application/list/"

type ApplicationList struct {
	AppReplacemsg                 *string                                  `json:"app-replacemsg,omitempty"`
	Comment                       *string                                  `json:"comment,omitempty"`
	ControlDefaultNetworkServices *string                                  `json:"control-default-network-services,omitempty"`
	DeepAppInspection             *string                                  `json:"deep-app-inspection,omitempty"`
	DefaultNetworkServices        *[]ApplicationListDefaultNetworkServices `json:"default-network-services,omitempty"`
	EnforceDefaultAppPort         *string                                  `json:"enforce-default-app-port,omitempty"`
	Entries                       *[]ApplicationListEntries                `json:"entries,omitempty"`
	ExtendedLog                   *string                                  `json:"extended-log,omitempty"`
	ForceInclusionSslDiSigs       *string                                  `json:"force-inclusion-ssl-di-sigs,omitempty"`
	Name                          *string                                  `json:"name,omitempty"`
	Options                       *string                                  `json:"options,omitempty"`
	OtherApplicationAction        *string                                  `json:"other-application-action,omitempty"`
	OtherApplicationLog           *string                                  `json:"other-application-log,omitempty"`
	P2pBlackList                  *string                                  `json:"p2p-black-list,omitempty"`
	P2pBlockList                  *string                                  `json:"p2p-block-list,omitempty"`
	ReplacemsgGroup               *string                                  `json:"replacemsg-group,omitempty"`
	UnknownApplicationAction      *string                                  `json:"unknown-application-action,omitempty"`
	UnknownApplicationLog         *string                                  `json:"unknown-application-log,omitempty"`
}

type ApplicationListDefaultNetworkServices struct {
	Id              *int64  `json:"id,omitempty"`
	Port            *int64  `json:"port,omitempty"`
	Services        *string `json:"services,omitempty"`
	ViolationAction *string `json:"violation-action,omitempty"`
}

type ApplicationListEntries struct {
	Action           *string                              `json:"action,omitempty"`
	Application      *[]ApplicationListEntriesApplication `json:"application,omitempty"`
	Behavior         *string                              `json:"behavior,omitempty"`
	Category         *[]ApplicationListEntriesCategory    `json:"category,omitempty"`
	Exclusion        *[]ApplicationListEntriesExclusion   `json:"exclusion,omitempty"`
	Id               *int64                               `json:"id,omitempty"`
	Log              *string                              `json:"log,omitempty"`
	LogPacket        *string                              `json:"log-packet,omitempty"`
	Parameters       *[]ApplicationListEntriesParameters  `json:"parameters,omitempty"`
	PerIpShaper      *string                              `json:"per-ip-shaper,omitempty"`
	Popularity       *string                              `json:"popularity,omitempty"`
	Protocols        *string                              `json:"protocols,omitempty"`
	Quarantine       *string                              `json:"quarantine,omitempty"`
	QuarantineExpiry *string                              `json:"quarantine-expiry,omitempty"`
	QuarantineLog    *string                              `json:"quarantine-log,omitempty"`
	RateCount        *int64                               `json:"rate-count,omitempty"`
	RateDuration     *int64                               `json:"rate-duration,omitempty"`
	RateMode         *string                              `json:"rate-mode,omitempty"`
	RateTrack        *string                              `json:"rate-track,omitempty"`
	Risk             *[]ApplicationListEntriesRisk        `json:"risk,omitempty"`
	SessionTtl       *int64                               `json:"session-ttl,omitempty"`
	Shaper           *string                              `json:"shaper,omitempty"`
	ShaperReverse    *string                              `json:"shaper-reverse,omitempty"`
	SubCategory      *[]ApplicationListEntriesSubCategory `json:"sub-category,omitempty"`
	Technology       *string                              `json:"technology,omitempty"`
	Vendor           *string                              `json:"vendor,omitempty"`
}

type ApplicationListEntriesApplication struct {
	Id *int64 `json:"id,omitempty"`
}

type ApplicationListEntriesCategory struct {
	Id *int64 `json:"id,omitempty"`
}

type ApplicationListEntriesExclusion struct {
	Id *int64 `json:"id,omitempty"`
}

type ApplicationListEntriesParameters struct {
	Id      *int64                                     `json:"id,omitempty"`
	Members *[]ApplicationListEntriesParametersMembers `json:"members,omitempty"`
	Value   *string                                    `json:"value,omitempty"`
}

type ApplicationListEntriesParametersMembers struct {
	Id    *int64  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ApplicationListEntriesRisk struct {
	Level *int64 `json:"level,omitempty"`
}

type ApplicationListEntriesSubCategory struct {
	Id *int64 `json:"id,omitempty"`
}
