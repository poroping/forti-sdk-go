package models

const SwitchControllerQosQosPolicyPath = "switch-controller.qos/qos-policy/"

type SwitchControllerQosQosPolicy struct {
	DefaultCos     *int64  `json:"default-cos,omitempty"`
	Name           *string `json:"name,omitempty"`
	QueuePolicy    *string `json:"queue-policy,omitempty"`
	TrustDot1pMap  *string `json:"trust-dot1p-map,omitempty"`
	TrustIpDscpMap *string `json:"trust-ip-dscp-map,omitempty"`
}
