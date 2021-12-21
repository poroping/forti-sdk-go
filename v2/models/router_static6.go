package models

const RouterStatic6Path = "router/static6/"

type RouterStatic6 struct {
	Bfd               *string                   `json:"bfd,omitempty"`
	Blackhole         *string                   `json:"blackhole,omitempty"`
	Comment           *string                   `json:"comment,omitempty"`
	Device            *string                   `json:"device,omitempty"`
	Devindex          *int64                    `json:"devindex,omitempty"`
	Distance          *int64                    `json:"distance,omitempty"`
	Dst               *string                   `json:"dst,omitempty"`
	DynamicGateway    *string                   `json:"dynamic-gateway,omitempty"`
	Gateway           *string                   `json:"gateway,omitempty"`
	LinkMonitorExempt *string                   `json:"link-monitor-exempt,omitempty"`
	Priority          *int64                    `json:"priority,omitempty"`
	Sdwan             *string                   `json:"sdwan,omitempty"`
	SdwanZone         *[]RouterStatic6SdwanZone `json:"sdwan-zone,omitempty"`
	SeqNum            *int64                    `json:"seq-num,omitempty"`
	Status            *string                   `json:"status,omitempty"`
	VirtualWanLink    *string                   `json:"virtual-wan-link,omitempty"`
	Vrf               *int64                    `json:"vrf,omitempty"`
}

type RouterStatic6SdwanZone struct {
	Name *string `json:"name,omitempty"`
}
