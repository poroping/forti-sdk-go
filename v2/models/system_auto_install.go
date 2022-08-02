package models

const SystemAutoInstallPath = "system/auto-install/"

type SystemAutoInstall struct {
	AutoInstallConfig *string `json:"auto-install-config,omitempty"`
	AutoInstallImage  *string `json:"auto-install-image,omitempty"`
	DefaultConfigFile *string `json:"default-config-file,omitempty"`
	DefaultImageFile  *string `json:"default-image-file,omitempty"`
}
