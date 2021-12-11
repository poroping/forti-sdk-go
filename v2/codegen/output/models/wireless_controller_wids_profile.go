package models

const WirelessControllerWidsProfilePath = "wireless-controller/wids-profile/"

type WirelessControllerWidsProfile struct {
	ApAutoSuppress           *string                                                 `json:"ap-auto-suppress,omitempty"`
	ApBgscanDisableSchedules []WirelessControllerWidsProfileApBgscanDisableSchedules `json:"ap-bgscan-disable-schedules,omitempty"`
	ApBgscanDuration         *int64                                                  `json:"ap-bgscan-duration,omitempty"`
	ApBgscanIdle             *int64                                                  `json:"ap-bgscan-idle,omitempty"`
	ApBgscanIntv             *int64                                                  `json:"ap-bgscan-intv,omitempty"`
	ApBgscanPeriod           *int64                                                  `json:"ap-bgscan-period,omitempty"`
	ApBgscanReportIntv       *int64                                                  `json:"ap-bgscan-report-intv,omitempty"`
	ApFgscanReportIntv       *int64                                                  `json:"ap-fgscan-report-intv,omitempty"`
	ApScan                   *string                                                 `json:"ap-scan,omitempty"`
	ApScanPassive            *string                                                 `json:"ap-scan-passive,omitempty"`
	ApScanThreshold          *string                                                 `json:"ap-scan-threshold,omitempty"`
	AsleapAttack             *string                                                 `json:"asleap-attack,omitempty"`
	AssocFloodThresh         *int64                                                  `json:"assoc-flood-thresh,omitempty"`
	AssocFloodTime           *int64                                                  `json:"assoc-flood-time,omitempty"`
	AssocFrameFlood          *string                                                 `json:"assoc-frame-flood,omitempty"`
	AuthFloodThresh          *int64                                                  `json:"auth-flood-thresh,omitempty"`
	AuthFloodTime            *int64                                                  `json:"auth-flood-time,omitempty"`
	AuthFrameFlood           *string                                                 `json:"auth-frame-flood,omitempty"`
	Comment                  *string                                                 `json:"comment,omitempty"`
	DeauthBroadcast          *string                                                 `json:"deauth-broadcast,omitempty"`
	DeauthUnknownSrcThresh   *int64                                                  `json:"deauth-unknown-src-thresh,omitempty"`
	EapolFailFlood           *string                                                 `json:"eapol-fail-flood,omitempty"`
	EapolFailIntv            *int64                                                  `json:"eapol-fail-intv,omitempty"`
	EapolFailThresh          *int64                                                  `json:"eapol-fail-thresh,omitempty"`
	EapolLogoffFlood         *string                                                 `json:"eapol-logoff-flood,omitempty"`
	EapolLogoffIntv          *int64                                                  `json:"eapol-logoff-intv,omitempty"`
	EapolLogoffThresh        *int64                                                  `json:"eapol-logoff-thresh,omitempty"`
	EapolPreFailFlood        *string                                                 `json:"eapol-pre-fail-flood,omitempty"`
	EapolPreFailIntv         *int64                                                  `json:"eapol-pre-fail-intv,omitempty"`
	EapolPreFailThresh       *int64                                                  `json:"eapol-pre-fail-thresh,omitempty"`
	EapolPreSuccFlood        *string                                                 `json:"eapol-pre-succ-flood,omitempty"`
	EapolPreSuccIntv         *int64                                                  `json:"eapol-pre-succ-intv,omitempty"`
	EapolPreSuccThresh       *int64                                                  `json:"eapol-pre-succ-thresh,omitempty"`
	EapolStartFlood          *string                                                 `json:"eapol-start-flood,omitempty"`
	EapolStartIntv           *int64                                                  `json:"eapol-start-intv,omitempty"`
	EapolStartThresh         *int64                                                  `json:"eapol-start-thresh,omitempty"`
	EapolSuccFlood           *string                                                 `json:"eapol-succ-flood,omitempty"`
	EapolSuccIntv            *int64                                                  `json:"eapol-succ-intv,omitempty"`
	EapolSuccThresh          *int64                                                  `json:"eapol-succ-thresh,omitempty"`
	InvalidMacOui            *string                                                 `json:"invalid-mac-oui,omitempty"`
	LongDurationAttack       *string                                                 `json:"long-duration-attack,omitempty"`
	LongDurationThresh       *int64                                                  `json:"long-duration-thresh,omitempty"`
	Name                     *string                                                 `json:"name,omitempty"`
	NullSsidProbeResp        *string                                                 `json:"null-ssid-probe-resp,omitempty"`
	SensorMode               *string                                                 `json:"sensor-mode,omitempty"`
	SpoofedDeauth            *string                                                 `json:"spoofed-deauth,omitempty"`
	WeakWepIv                *string                                                 `json:"weak-wep-iv,omitempty"`
	WirelessBridge           *string                                                 `json:"wireless-bridge,omitempty"`
}

type WirelessControllerWidsProfileApBgscanDisableSchedules struct {
	Name *string `json:"name,omitempty"`
}
