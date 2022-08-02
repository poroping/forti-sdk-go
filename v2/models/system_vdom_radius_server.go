package models

const SystemVdomRadiusServerPath = "system/vdom-radius-server/"

type SystemVdomRadiusServer struct {
	Name             *string `json:"name,omitempty"`
	RadiusServerVdom *string `json:"radius-server-vdom,omitempty"`
	Status           *string `json:"status,omitempty"`
}
