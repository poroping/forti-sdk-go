package models

const RouterPolicyPath = "router/policy/"

type RouterPolicy struct {
	Action                *string                              `json:"action,omitempty"`
	Comments              *string                              `json:"comments,omitempty"`
	Dst                   *[]RouterPolicyDst                   `json:"dst,omitempty"`
	DstNegate             *string                              `json:"dst-negate,omitempty"`
	Dstaddr               *[]RouterPolicyDstaddr               `json:"dstaddr,omitempty"`
	EndPort               *int64                               `json:"end-port,omitempty"`
	EndSourcePort         *int64                               `json:"end-source-port,omitempty"`
	Gateway               *string                              `json:"gateway,omitempty"`
	InputDevice           *[]RouterPolicyInputDevice           `json:"input-device,omitempty"`
	InputDeviceNegate     *string                              `json:"input-device-negate,omitempty"`
	InternetServiceCustom *[]RouterPolicyInternetServiceCustom `json:"internet-service-custom,omitempty"`
	InternetServiceId     *[]RouterPolicyInternetServiceId     `json:"internet-service-id,omitempty"`
	OutputDevice          *string                              `json:"output-device,omitempty"`
	Protocol              *int64                               `json:"protocol,omitempty"`
	SeqNum                *int64                               `json:"seq-num,omitempty"`
	Src                   *[]RouterPolicySrc                   `json:"src,omitempty"`
	SrcNegate             *string                              `json:"src-negate,omitempty"`
	Srcaddr               *[]RouterPolicySrcaddr               `json:"srcaddr,omitempty"`
	StartPort             *int64                               `json:"start-port,omitempty"`
	StartSourcePort       *int64                               `json:"start-source-port,omitempty"`
	Status                *string                              `json:"status,omitempty"`
	Tos                   *string                              `json:"tos,omitempty"`
	TosMask               *string                              `json:"tos-mask,omitempty"`
}

const RouterPolicyDstPath = "router/policy/dst/"

type RouterPolicyDst struct {
	Subnet *string `json:"subnet,omitempty"`
}

const RouterPolicyDstaddrPath = "router/policy/dstaddr/"

type RouterPolicyDstaddr struct {
	Name *string `json:"name,omitempty"`
}

const RouterPolicyInputDevicePath = "router/policy/input-device/"

type RouterPolicyInputDevice struct {
	Name *string `json:"name,omitempty"`
}

const RouterPolicyInternetServiceCustomPath = "router/policy/internet-service-custom/"

type RouterPolicyInternetServiceCustom struct {
	Name *string `json:"name,omitempty"`
}

const RouterPolicyInternetServiceIdPath = "router/policy/internet-service-id/"

type RouterPolicyInternetServiceId struct {
	Id *int64 `json:"id,omitempty"`
}

const RouterPolicySrcPath = "router/policy/src/"

type RouterPolicySrc struct {
	Subnet *string `json:"subnet,omitempty"`
}

const RouterPolicySrcaddrPath = "router/policy/srcaddr/"

type RouterPolicySrcaddr struct {
	Name *string `json:"name,omitempty"`
}
