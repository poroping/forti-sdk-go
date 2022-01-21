package models

const VpnSslWebHostCheckSoftwarePath = "vpn.ssl.web/host-check-software/"

type VpnSslWebHostCheckSoftware struct {
	CheckItemList *[]VpnSslWebHostCheckSoftwareCheckItemList `json:"check-item-list,omitempty"`
	Guid          *string                                    `json:"guid,omitempty"`
	Name          *string                                    `json:"name,omitempty"`
	OsType        *string                                    `json:"os-type,omitempty"`
	Type          *string                                    `json:"type,omitempty"`
	Version       *string                                    `json:"version,omitempty"`
}

type VpnSslWebHostCheckSoftwareCheckItemList struct {
	Action  *string                                        `json:"action,omitempty"`
	Id      *int64                                         `json:"id,omitempty"`
	Md5s    *[]VpnSslWebHostCheckSoftwareCheckItemListMd5s `json:"md5s,omitempty"`
	Target  *string                                        `json:"target,omitempty"`
	Type    *string                                        `json:"type,omitempty"`
	Version *string                                        `json:"version,omitempty"`
}

type VpnSslWebHostCheckSoftwareCheckItemListMd5s struct {
	Id *string `json:"id,omitempty"`
}
