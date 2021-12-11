package models

const SystemSamlPath = "system/saml/"

type SystemSaml struct {
	BindingProtocol    *string                      `json:"binding-protocol,omitempty"`
	Cert               *string                      `json:"cert,omitempty"`
	DefaultLoginPage   *string                      `json:"default-login-page,omitempty"`
	DefaultProfile     *string                      `json:"default-profile,omitempty"`
	EntityId           *string                      `json:"entity-id,omitempty"`
	IdpCert            *string                      `json:"idp-cert,omitempty"`
	IdpEntityId        *string                      `json:"idp-entity-id,omitempty"`
	IdpSingleLogoutUrl *string                      `json:"idp-single-logout-url,omitempty"`
	IdpSingleSignOnUrl *string                      `json:"idp-single-sign-on-url,omitempty"`
	Life               *int64                       `json:"life,omitempty"`
	PortalUrl          *string                      `json:"portal-url,omitempty"`
	Role               *string                      `json:"role,omitempty"`
	ServerAddress      *string                      `json:"server-address,omitempty"`
	ServiceProviders   []SystemSamlServiceProviders `json:"service-providers,omitempty"`
	SingleLogoutUrl    *string                      `json:"single-logout-url,omitempty"`
	SingleSignOnUrl    *string                      `json:"single-sign-on-url,omitempty"`
	Status             *string                      `json:"status,omitempty"`
	Tolerance          *int64                       `json:"tolerance,omitempty"`
}

type SystemSamlServiceProviders struct {
	AssertionAttributes []SystemSamlServiceProvidersAssertionAttributes `json:"assertion-attributes,omitempty"`
	IdpEntityId         *string                                         `json:"idp-entity-id,omitempty"`
	IdpSingleLogoutUrl  *string                                         `json:"idp-single-logout-url,omitempty"`
	IdpSingleSignOnUrl  *string                                         `json:"idp-single-sign-on-url,omitempty"`
	Name                *string                                         `json:"name,omitempty"`
	Prefix              *string                                         `json:"prefix,omitempty"`
	SpBindingProtocol   *string                                         `json:"sp-binding-protocol,omitempty"`
	SpCert              *string                                         `json:"sp-cert,omitempty"`
	SpEntityId          *string                                         `json:"sp-entity-id,omitempty"`
	SpPortalUrl         *string                                         `json:"sp-portal-url,omitempty"`
	SpSingleLogoutUrl   *string                                         `json:"sp-single-logout-url,omitempty"`
	SpSingleSignOnUrl   *string                                         `json:"sp-single-sign-on-url,omitempty"`
}

type SystemSamlServiceProvidersAssertionAttributes struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}
