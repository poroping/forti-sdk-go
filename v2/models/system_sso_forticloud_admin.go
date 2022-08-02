package models

const SystemSsoForticloudAdminPath = "system/sso-forticloud-admin/"

type SystemSsoForticloudAdmin struct {
	Name *string                         `json:"name,omitempty"`
	Vdom *[]SystemSsoForticloudAdminVdom `json:"vdom,omitempty"`
}

const SystemSsoForticloudAdminVdomPath = "system/sso-forticloud-admin/vdom/"

type SystemSsoForticloudAdminVdom struct {
	Name *string `json:"name,omitempty"`
}
