package models

const RouterPolicy6Path = "router/policy6/"

type RouterPolicy6 struct {
	Comments     *string                     `json:"comments,omitempty"`
	Dst          *string                     `json:"dst,omitempty"`
	EndPort      *float64                    `json:"end-port,omitempty"`
	Gateway      *string                     `json:"gateway,omitempty"`
	InputDevice  *[]RouterPolicy6InputDevice `json:"input-device,omitempty"`
	OutputDevice *string                     `json:"output-device,omitempty"`
	Protocol     *float64                    `json:"protocol,omitempty"`
	SeqNum       *float64                    `json:"seq-num,omitempty"`
	Src          *string                     `json:"src,omitempty"`
	StartPort    *float64                    `json:"start-port,omitempty"`
	Status       *string                     `json:"status,omitempty"`
	Tos          *string                     `json:"tos,omitempty"`
	TosMask      *string                     `json:"tos-mask,omitempty"`
}

type RouterPolicy6InputDevice struct {
	Name *string `json:"name,omitempty"`
}
