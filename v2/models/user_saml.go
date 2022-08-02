package models

const UserSamlPath = "user/saml/"

type UserSaml struct {
	AdfsClaim          *string `json:"adfs-claim,omitempty"`
	Cert               *string `json:"cert,omitempty"`
	ClockTolerance     *int64  `json:"clock-tolerance,omitempty"`
	DigestMethod       *string `json:"digest-method,omitempty"`
	EntityId           *string `json:"entity-id,omitempty"`
	GroupClaimType     *string `json:"group-claim-type,omitempty"`
	GroupName          *string `json:"group-name,omitempty"`
	IdpCert            *string `json:"idp-cert,omitempty"`
	IdpEntityId        *string `json:"idp-entity-id,omitempty"`
	IdpSingleLogoutUrl *string `json:"idp-single-logout-url,omitempty"`
	IdpSingleSignOnUrl *string `json:"idp-single-sign-on-url,omitempty"`
	LimitRelaystate    *string `json:"limit-relaystate,omitempty"`
	Name               *string `json:"name,omitempty"`
	SingleLogoutUrl    *string `json:"single-logout-url,omitempty"`
	SingleSignOnUrl    *string `json:"single-sign-on-url,omitempty"`
	UserClaimType      *string `json:"user-claim-type,omitempty"`
	UserName           *string `json:"user-name,omitempty"`
}
