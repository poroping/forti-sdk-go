package models

const SystemVirtualSwitchPath = "system/virtual-switch/"

type SystemVirtualSwitch struct {
	Name           *string                   `json:"name,omitempty"`
	PhysicalSwitch *string                   `json:"physical-switch,omitempty"`
	Port           []SystemVirtualSwitchPort `json:"port,omitempty"`
	Span           *string                   `json:"span,omitempty"`
	SpanDestPort   *string                   `json:"span-dest-port,omitempty"`
	SpanDirection  *string                   `json:"span-direction,omitempty"`
	SpanSourcePort *string                   `json:"span-source-port,omitempty"`
	Vlan           *int64                    `json:"vlan,omitempty"`
}

type SystemVirtualSwitchPort struct {
	Alias *string `json:"alias,omitempty"`
	Name  *string `json:"name,omitempty"`
}
