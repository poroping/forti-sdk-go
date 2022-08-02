package models

const SystemSwitchInterfacePath = "system/switch-interface/"

type SystemSwitchInterface struct {
	IntraSwitchPolicy *string                                `json:"intra-switch-policy,omitempty"`
	MacTtl            *int64                                 `json:"mac-ttl,omitempty"`
	Member            *[]SystemSwitchInterfaceMember         `json:"member,omitempty"`
	Name              *string                                `json:"name,omitempty"`
	Span              *string                                `json:"span,omitempty"`
	SpanDestPort      *string                                `json:"span-dest-port,omitempty"`
	SpanDirection     *string                                `json:"span-direction,omitempty"`
	SpanSourcePort    *[]SystemSwitchInterfaceSpanSourcePort `json:"span-source-port,omitempty"`
	Type              *string                                `json:"type,omitempty"`
	Vdom              *string                                `json:"vdom,omitempty"`
}

const SystemSwitchInterfaceMemberPath = "system/switch-interface/member/"

type SystemSwitchInterfaceMember struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}

const SystemSwitchInterfaceSpanSourcePortPath = "system/switch-interface/span-source-port/"

type SystemSwitchInterfaceSpanSourcePort struct {
	InterfaceName *string `json:"interface-name,omitempty"`
}
