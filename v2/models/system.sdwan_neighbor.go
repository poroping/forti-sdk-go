package models

const SystemsdwanNeighborPath = "system.sdwan/neighbor/"

type SystemsdwanNeighbor struct {
	HealthCheck *string  `json:"health-check,omitempty"`
	Ip          *string  `json:"ip,omitempty"`
	Member      *float64 `json:"member,omitempty"`
	Mode        *string  `json:"mode,omitempty"`
	Role        *string  `json:"role,omitempty"`
	SlaId       *float64 `json:"sla-id,omitempty"`
}
