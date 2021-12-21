package models

const FirewallShapingProfilePath = "firewall/shaping-profile/"

type FirewallShapingProfile struct {
	Comment        *string                                 `json:"comment,omitempty"`
	DefaultClassId *float64                                `json:"default-class-id,omitempty"`
	ProfileName    *string                                 `json:"profile-name,omitempty"`
	ShapingEntries *[]FirewallShapingProfileShapingEntries `json:"shaping-entries,omitempty"`
	Type           *string                                 `json:"type,omitempty"`
}

type FirewallShapingProfileShapingEntries struct {
	BurstInMsec                   *float64 `json:"burst-in-msec,omitempty"`
	CburstInMsec                  *float64 `json:"cburst-in-msec,omitempty"`
	ClassId                       *float64 `json:"class-id,omitempty"`
	GuaranteedBandwidthPercentage *float64 `json:"guaranteed-bandwidth-percentage,omitempty"`
	Id                            *float64 `json:"id,omitempty"`
	Limit                         *float64 `json:"limit,omitempty"`
	Max                           *float64 `json:"max,omitempty"`
	MaximumBandwidthPercentage    *float64 `json:"maximum-bandwidth-percentage,omitempty"`
	Min                           *float64 `json:"min,omitempty"`
	Priority                      *string  `json:"priority,omitempty"`
	RedProbability                *float64 `json:"red-probability,omitempty"`
}
