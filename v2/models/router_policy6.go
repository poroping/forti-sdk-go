package models

const RouterPolicy6Path = "router/policy6/"

type RouterPolicy6 struct {
	Comments     *string                     `json:"comments,omitempty"`
	Dst          *string                     `json:"dst,omitempty"`
	EndPort      *int64                      `json:"end-port,omitempty"`
	Gateway      *string                     `json:"gateway,omitempty"`
	InputDevice  *[]RouterPolicy6InputDevice `json:"input-device,omitempty"`
	OutputDevice *string                     `json:"output-device,omitempty"`
	Protocol     *int64                      `json:"protocol,omitempty"`
	SeqNum       *int64                      `json:"seq-num,omitempty"`
	Src          *string                     `json:"src,omitempty"`
	StartPort    *int64                      `json:"start-port,omitempty"`
	Status       *string                     `json:"status,omitempty"`
	Tos          *string                     `json:"tos,omitempty"`
	TosMask      *string                     `json:"tos-mask,omitempty"`
}

type RouterPolicy6InputDevice struct {
	Name *string `json:"name,omitempty"`
}
