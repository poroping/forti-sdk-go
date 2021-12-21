package models

const RouterPolicyPath = "router/policy/"

type RouterPolicy struct {
	Action                *string                              `json:"action,omitempty"`
	Comments              *string                              `json:"comments,omitempty"`
	Dst                   *[]RouterPolicyDst                   `json:"dst,omitempty"`
	DstNegate             *string                              `json:"dst-negate,omitempty"`
	Dstaddr               *[]RouterPolicyDstaddr               `json:"dstaddr,omitempty"`
	EndPort               *float64                             `json:"end-port,omitempty"`
	EndSourcePort         *float64                             `json:"end-source-port,omitempty"`
	Gateway               *string                              `json:"gateway,omitempty"`
	InputDevice           *[]RouterPolicyInputDevice           `json:"input-device,omitempty"`
	InputDeviceNegate     *string                              `json:"input-device-negate,omitempty"`
	InternetServiceCustom *[]RouterPolicyInternetServiceCustom `json:"internet-service-custom,omitempty"`
	InternetServiceId     *[]RouterPolicyInternetServiceId     `json:"internet-service-id,omitempty"`
	OutputDevice          *string                              `json:"output-device,omitempty"`
	Protocol              *float64                             `json:"protocol,omitempty"`
	SeqNum                *float64                             `json:"seq-num,omitempty"`
	Src                   *[]RouterPolicySrc                   `json:"src,omitempty"`
	SrcNegate             *string                              `json:"src-negate,omitempty"`
	Srcaddr               *[]RouterPolicySrcaddr               `json:"srcaddr,omitempty"`
	StartPort             *float64                             `json:"start-port,omitempty"`
	StartSourcePort       *float64                             `json:"start-source-port,omitempty"`
	Status                *string                              `json:"status,omitempty"`
	Tos                   *string                              `json:"tos,omitempty"`
	TosMask               *string                              `json:"tos-mask,omitempty"`
}

type RouterPolicyDst struct {
	Subnet *string `json:"subnet,omitempty"`
}

type RouterPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

type RouterPolicyInputDevice struct {
	Name *string `json:"name,omitempty"`
}

type RouterPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

type RouterPolicyInternetServiceId struct {
	Id *float64 `json:"id,omitempty"`
}

type RouterPolicySrc struct {
	Subnet *string `json:"subnet,omitempty"`
}

type RouterPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
