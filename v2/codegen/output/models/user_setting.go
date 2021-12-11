package models

const UserSettingPath = "user/setting/"

type UserSetting struct {
	AuthBlackoutTime          *int64                 `json:"auth-blackout-time,omitempty"`
	AuthCaCert                *string                `json:"auth-ca-cert,omitempty"`
	AuthCert                  *string                `json:"auth-cert,omitempty"`
	AuthHttpBasic             *string                `json:"auth-http-basic,omitempty"`
	AuthInvalidMax            *int64                 `json:"auth-invalid-max,omitempty"`
	AuthLockoutDuration       *int64                 `json:"auth-lockout-duration,omitempty"`
	AuthLockoutThreshold      *int64                 `json:"auth-lockout-threshold,omitempty"`
	AuthOnDemand              *string                `json:"auth-on-demand,omitempty"`
	AuthPortalTimeout         *int64                 `json:"auth-portal-timeout,omitempty"`
	AuthPorts                 []UserSettingAuthPorts `json:"auth-ports,omitempty"`
	AuthSecureHttp            *string                `json:"auth-secure-http,omitempty"`
	AuthSrcMac                *string                `json:"auth-src-mac,omitempty"`
	AuthSslAllowRenegotiation *string                `json:"auth-ssl-allow-renegotiation,omitempty"`
	AuthSslMaxProtoVersion    *string                `json:"auth-ssl-max-proto-version,omitempty"`
	AuthSslMinProtoVersion    *string                `json:"auth-ssl-min-proto-version,omitempty"`
	AuthSslSigalgs            *string                `json:"auth-ssl-sigalgs,omitempty"`
	AuthTimeout               *int64                 `json:"auth-timeout,omitempty"`
	AuthTimeoutType           *string                `json:"auth-timeout-type,omitempty"`
	AuthType                  *string                `json:"auth-type,omitempty"`
	PerPolicyDisclaimer       *string                `json:"per-policy-disclaimer,omitempty"`
	RadiusSesTimeoutAct       *string                `json:"radius-ses-timeout-act,omitempty"`
}

type UserSettingAuthPorts struct {
	Id   *int64  `json:"id,omitempty"`
	Port *int64  `json:"port,omitempty"`
	Type *string `json:"type,omitempty"`
}
