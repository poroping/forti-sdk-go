package models

const SystemSsoAdminPath = "system/sso-admin/"

type SystemSsoAdmin struct {
	Accprofile *string               `json:"accprofile,omitempty"`
	Name       *string               `json:"name,omitempty"`
	Vdom       *[]SystemSsoAdminVdom `json:"vdom,omitempty"`
}

const SystemSsoAdminVdomPath = "system/sso-admin/vdom/"

type SystemSsoAdminVdom struct {
	Name *string `json:"name,omitempty"`
}
