package models

const VpnsslwebUserGroupBookmarkPath = "vpn/ssl/web/user-group-bookmark/"

type VpnsslwebUserGroupBookmark struct {
	Bookmarks *[]VpnsslwebUserGroupBookmarkBookmarks `json:"bookmarks,omitempty"`
	Name      *string                                `json:"name,omitempty"`
}

type VpnsslwebUserGroupBookmarkBookmarks struct {
	AdditionalParams      *string                                        `json:"additional-params,omitempty"`
	Apptype               *string                                        `json:"apptype,omitempty"`
	ColorDepth            *string                                        `json:"color-depth,omitempty"`
	Description           *string                                        `json:"description,omitempty"`
	Domain                *string                                        `json:"domain,omitempty"`
	Folder                *string                                        `json:"folder,omitempty"`
	FormData              *[]VpnsslwebUserGroupBookmarkBookmarksFormData `json:"form-data,omitempty"`
	Host                  *string                                        `json:"host,omitempty"`
	KeyboardLayout        *string                                        `json:"keyboard-layout,omitempty"`
	ListeningPort         *int64                                         `json:"listening-port,omitempty"`
	LoadBalancingInfo     *string                                        `json:"load-balancing-info,omitempty"`
	LogonPassword         *string                                        `json:"logon-password,omitempty"`
	LogonUser             *string                                        `json:"logon-user,omitempty"`
	Name                  *string                                        `json:"name,omitempty"`
	Port                  *int64                                         `json:"port,omitempty"`
	PreconnectionBlob     *string                                        `json:"preconnection-blob,omitempty"`
	PreconnectionId       *int64                                         `json:"preconnection-id,omitempty"`
	RemotePort            *int64                                         `json:"remote-port,omitempty"`
	RestrictedAdmin       *string                                        `json:"restricted-admin,omitempty"`
	Security              *string                                        `json:"security,omitempty"`
	SendPreconnectionId   *string                                        `json:"send-preconnection-id,omitempty"`
	ServerLayout          *string                                        `json:"server-layout,omitempty"`
	ShowStatusWindow      *string                                        `json:"show-status-window,omitempty"`
	Sso                   *string                                        `json:"sso,omitempty"`
	SsoCredential         *string                                        `json:"sso-credential,omitempty"`
	SsoCredentialSentOnce *string                                        `json:"sso-credential-sent-once,omitempty"`
	SsoPassword           *string                                        `json:"sso-password,omitempty"`
	SsoUsername           *string                                        `json:"sso-username,omitempty"`
	Url                   *string                                        `json:"url,omitempty"`
}

type VpnsslwebUserGroupBookmarkBookmarksFormData struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
