package models

const SystemsdwanZonePath = "system/sdwan/zone/"

type SystemsdwanZone struct {
	Name               *string `json:"name,omitempty"`
	ServiceSlaTieBreak *string `json:"service-sla-tie-break,omitempty"`
}
