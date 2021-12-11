package models

const VpnsslwebHostCheckSoftwarePath = "vpn.ssl.web/host-check-software/"

type VpnsslwebHostCheckSoftware struct {
	CheckItemList []VpnsslwebHostCheckSoftwareCheckItemList `json:"check-item-list,omitempty"`
	Guid          *string                                   `json:"guid,omitempty"`
	Name          *string                                   `json:"name,omitempty"`
	OsType        *string                                   `json:"os-type,omitempty"`
	Type          *string                                   `json:"type,omitempty"`
	Version       *string                                   `json:"version,omitempty"`
}

type VpnsslwebHostCheckSoftwareCheckItemList struct {
	Action  *string                                       `json:"action,omitempty"`
	Id      *int64                                        `json:"id,omitempty"`
	Md5s    []VpnsslwebHostCheckSoftwareCheckItemListMd5s `json:"md5s,omitempty"`
	Target  *string                                       `json:"target,omitempty"`
	Type    *string                                       `json:"type,omitempty"`
	Version *string                                       `json:"version,omitempty"`
}

type VpnsslwebHostCheckSoftwareCheckItemListMd5s struct {
	Id *string `json:"id,omitempty"`
}
