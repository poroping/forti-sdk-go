package models

const SystemNpuPath = "system/npu/"

type SystemNpu struct {
	CapwapOffload               *string `json:"capwap-offload,omitempty"`
	DedicatedManagementAffinity *string `json:"dedicated-management-affinity,omitempty"`
	DedicatedManagementCpu      *string `json:"dedicated-management-cpu,omitempty"`
	IpsecMtuOverride            *string `json:"ipsec-mtu-override,omitempty"`
}
