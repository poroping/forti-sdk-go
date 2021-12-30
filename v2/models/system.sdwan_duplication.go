package models

const SystemsdwanDuplicationPath = "system/sdwan/duplication/"

type SystemsdwanDuplication struct {
	Dstaddr             *[]SystemsdwanDuplicationDstaddr   `json:"dstaddr,omitempty"`
	Dstaddr6            *[]SystemsdwanDuplicationDstaddr6  `json:"dstaddr6,omitempty"`
	Dstintf             *[]SystemsdwanDuplicationDstintf   `json:"dstintf,omitempty"`
	Fosid               *int64                             `json:"fosid,omitempty"`
	PacketDeDuplication *string                            `json:"packet-de-duplication,omitempty"`
	PacketDuplication   *string                            `json:"packet-duplication,omitempty"`
	Service             *[]SystemsdwanDuplicationService   `json:"service,omitempty"`
	ServiceId           *[]SystemsdwanDuplicationServiceId `json:"service-id,omitempty"`
	Srcaddr             *[]SystemsdwanDuplicationSrcaddr   `json:"srcaddr,omitempty"`
	Srcaddr6            *[]SystemsdwanDuplicationSrcaddr6  `json:"srcaddr6,omitempty"`
	Srcintf             *[]SystemsdwanDuplicationSrcintf   `json:"srcintf,omitempty"`
}

type SystemsdwanDuplicationDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanDuplicationDstaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanDuplicationDstintf struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanDuplicationService struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanDuplicationServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemsdwanDuplicationSrcaddr struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanDuplicationSrcaddr6 struct {
	Name *string `json:"name,omitempty"`
}

type SystemsdwanDuplicationSrcintf struct {
	Name *string `json:"name,omitempty"`
}
