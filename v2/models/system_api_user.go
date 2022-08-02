package models

const SystemApiUserPath = "system/api-user/"

type SystemApiUser struct {
	Accprofile      *string                   `json:"accprofile,omitempty"`
	ApiKey          *string                   `json:"api-key,omitempty"`
	Comments        *string                   `json:"comments,omitempty"`
	CorsAllowOrigin *string                   `json:"cors-allow-origin,omitempty"`
	Name            *string                   `json:"name,omitempty"`
	PeerAuth        *string                   `json:"peer-auth,omitempty"`
	PeerGroup       *string                   `json:"peer-group,omitempty"`
	Schedule        *string                   `json:"schedule,omitempty"`
	Trusthost       *[]SystemApiUserTrusthost `json:"trusthost,omitempty"`
	Vdom            *[]SystemApiUserVdom      `json:"vdom,omitempty"`
}

const SystemApiUserTrusthostPath = "system/api-user/trusthost/"

type SystemApiUserTrusthost struct {
	Id            *int64  `json:"id,omitempty"`
	Ipv4Trusthost *string `json:"ipv4-trusthost,omitempty"`
	Ipv6Trusthost *string `json:"ipv6-trusthost,omitempty"`
	Type          *string `json:"type,omitempty"`
}

const SystemApiUserVdomPath = "system/api-user/vdom/"

type SystemApiUserVdom struct {
	Name *string `json:"name,omitempty"`
}
