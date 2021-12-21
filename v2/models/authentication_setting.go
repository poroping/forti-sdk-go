package models

const AuthenticationSettingPath = "authentication/setting/"

type AuthenticationSetting struct {
	ActiveAuthScheme      *string                            `json:"active-auth-scheme,omitempty"`
	AuthHttps             *string                            `json:"auth-https,omitempty"`
	CaptivePortal         *string                            `json:"captive-portal,omitempty"`
	CaptivePortalIp       *string                            `json:"captive-portal-ip,omitempty"`
	CaptivePortalIp6      *string                            `json:"captive-portal-ip6,omitempty"`
	CaptivePortalPort     *float64                           `json:"captive-portal-port,omitempty"`
	CaptivePortalSslPort  *float64                           `json:"captive-portal-ssl-port,omitempty"`
	CaptivePortalType     *string                            `json:"captive-portal-type,omitempty"`
	CaptivePortal6        *string                            `json:"captive-portal6,omitempty"`
	CertAuth              *string                            `json:"cert-auth,omitempty"`
	CertCaptivePortal     *string                            `json:"cert-captive-portal,omitempty"`
	CertCaptivePortalIp   *string                            `json:"cert-captive-portal-ip,omitempty"`
	CertCaptivePortalPort *float64                           `json:"cert-captive-portal-port,omitempty"`
	DevRange              *[]AuthenticationSettingDevRange   `json:"dev-range,omitempty"`
	SsoAuthScheme         *string                            `json:"sso-auth-scheme,omitempty"`
	UserCertCa            *[]AuthenticationSettingUserCertCa `json:"user-cert-ca,omitempty"`
}

type AuthenticationSettingDevRange struct {
	Name *string `json:"name,omitempty"`
}

type AuthenticationSettingUserCertCa struct {
	Name *string `json:"name,omitempty"`
}
