package models

const RouterStaticPath = "router/static/"

type RouterStatic struct {
	Bfd                   *string                  `json:"bfd,omitempty"`
	Blackhole             *string                  `json:"blackhole,omitempty"`
	Comment               *string                  `json:"comment,omitempty"`
	Device                *string                  `json:"device,omitempty"`
	Distance              *float64                 `json:"distance,omitempty"`
	Dst                   *string                  `json:"dst,omitempty"`
	Dstaddr               *string                  `json:"dstaddr,omitempty"`
	DynamicGateway        *string                  `json:"dynamic-gateway,omitempty"`
	Gateway               *string                  `json:"gateway,omitempty"`
	InternetService       *float64                 `json:"internet-service,omitempty"`
	InternetServiceCustom *string                  `json:"internet-service-custom,omitempty"`
	LinkMonitorExempt     *string                  `json:"link-monitor-exempt,omitempty"`
	Priority              *float64                 `json:"priority,omitempty"`
	Sdwan                 *string                  `json:"sdwan,omitempty"`
	SdwanZone             *[]RouterStaticSdwanZone `json:"sdwan-zone,omitempty"`
	SeqNum                *float64                 `json:"seq-num,omitempty"`
	Src                   *string                  `json:"src,omitempty"`
	Status                *string                  `json:"status,omitempty"`
	VirtualWanLink        *string                  `json:"virtual-wan-link,omitempty"`
	Vrf                   *float64                 `json:"vrf,omitempty"`
	Weight                *float64                 `json:"weight,omitempty"`
}

type RouterStaticSdwanZone struct {
	Name *string `json:"name,omitempty"`
}
