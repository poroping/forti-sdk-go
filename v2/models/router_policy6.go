package models

const RouterPolicy6Path = "router/policy6/"

type RouterPolicy6 struct {
	Action                *string                               `json:"action,omitempty"`
	Comments              *string                               `json:"comments,omitempty"`
	Dst                   *[]RouterPolicy6Dst                   `json:"dst,omitempty"`
	DstNegate             *string                               `json:"dst-negate,omitempty"`
	Dstaddr               *[]RouterPolicy6Dstaddr               `json:"dstaddr,omitempty"`
	EndPort               *int64                                `json:"end-port,omitempty"`
	Gateway               *string                               `json:"gateway,omitempty"`
	InputDevice           *[]RouterPolicy6InputDevice           `json:"input-device,omitempty"`
	InputDeviceNegate     *string                               `json:"input-device-negate,omitempty"`
	InternetServiceCustom *[]RouterPolicy6InternetServiceCustom `json:"internet-service-custom,omitempty"`
	InternetServiceId     *[]RouterPolicy6InternetServiceId     `json:"internet-service-id,omitempty"`
	OutputDevice          *string                               `json:"output-device,omitempty"`
	Protocol              *int64                                `json:"protocol,omitempty"`
	SeqNum                *int64                                `json:"seq-num,omitempty"`
	Src                   *[]RouterPolicy6Src                   `json:"src,omitempty"`
	SrcNegate             *string                               `json:"src-negate,omitempty"`
	Srcaddr               *[]RouterPolicy6Srcaddr               `json:"srcaddr,omitempty"`
	StartPort             *int64                                `json:"start-port,omitempty"`
	Status                *string                               `json:"status,omitempty"`
	Tos                   *string                               `json:"tos,omitempty"`
	TosMask               *string                               `json:"tos-mask,omitempty"`
}

type RouterPolicy6Dst struct {
	Addr6 *string `json:"addr6,omitempty"`
}

type RouterPolicy6Dstaddr struct {
	Name *string `json:"name,omitempty"`
}

type RouterPolicy6InputDevice struct {
	Name *string `json:"name,omitempty"`
}

type RouterPolicy6InternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type RouterPolicy6InternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

type RouterPolicy6Src struct {
	Addr6 *string `json:"addr6,omitempty"`
}

type RouterPolicy6Srcaddr struct {
	Name *string `json:"name,omitempty"`
}
