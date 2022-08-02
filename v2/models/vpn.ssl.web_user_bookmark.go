package models

const VpnSslWebUserBookmarkPath = "vpn.ssl.web/user-bookmark/"

type VpnSslWebUserBookmark struct {
	Bookmarks  *[]VpnSslWebUserBookmarkBookmarks `json:"bookmarks,omitempty"`
	CustomLang *string                           `json:"custom-lang,omitempty"`
	Name       *string                           `json:"name,omitempty"`
}

const VpnSslWebUserBookmarkBookmarksPath = "vpn.ssl.web/user-bookmark/bookmarks/"

type VpnSslWebUserBookmarkBookmarks struct {
	AdditionalParams      *string                                   `json:"additional-params,omitempty"`
	Apptype               *string                                   `json:"apptype,omitempty"`
	ColorDepth            *string                                   `json:"color-depth,omitempty"`
	Description           *string                                   `json:"description,omitempty"`
	Domain                *string                                   `json:"domain,omitempty"`
	Folder                *string                                   `json:"folder,omitempty"`
	FormData              *[]VpnSslWebUserBookmarkBookmarksFormData `json:"form-data,omitempty"`
	Height                *int64                                    `json:"height,omitempty"`
	Host                  *string                                   `json:"host,omitempty"`
	KeyboardLayout        *string                                   `json:"keyboard-layout,omitempty"`
	LoadBalancingInfo     *string                                   `json:"load-balancing-info,omitempty"`
	LogonPassword         *string                                   `json:"logon-password,omitempty"`
	LogonUser             *string                                   `json:"logon-user,omitempty"`
	Name                  *string                                   `json:"name,omitempty"`
	Port                  *int64                                    `json:"port,omitempty"`
	PreconnectionBlob     *string                                   `json:"preconnection-blob,omitempty"`
	PreconnectionId       *int64                                    `json:"preconnection-id,omitempty"`
	RestrictedAdmin       *string                                   `json:"restricted-admin,omitempty"`
	Security              *string                                   `json:"security,omitempty"`
	SendPreconnectionId   *string                                   `json:"send-preconnection-id,omitempty"`
	Sso                   *string                                   `json:"sso,omitempty"`
	SsoCredential         *string                                   `json:"sso-credential,omitempty"`
	SsoCredentialSentOnce *string                                   `json:"sso-credential-sent-once,omitempty"`
	SsoPassword           *string                                   `json:"sso-password,omitempty"`
	SsoUsername           *string                                   `json:"sso-username,omitempty"`
	Url                   *string                                   `json:"url,omitempty"`
	Width                 *int64                                    `json:"width,omitempty"`
}

const VpnSslWebUserBookmarkBookmarksFormDataPath = "vpn.ssl.web/user-bookmark/bookmarks/form-data/"

type VpnSslWebUserBookmarkBookmarksFormData struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
