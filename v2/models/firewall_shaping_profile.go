package models

const FirewallShapingProfilePath = "firewall/shaping-profile/"

type FirewallShapingProfile struct {
	Comment        *string                                `json:"comment,omitempty"`
	DefaultClassId *int64                                 `json:"default-class-id,omitempty"`
	ProfileName    *string                                `json:"profile-name,omitempty"`
	ShapingEntries []FirewallShapingProfileShapingEntries `json:"shaping-entries,omitempty"`
	Type           *string                                `json:"type,omitempty"`
}

type FirewallShapingProfileShapingEntries struct {
	BurstInMsec                   *int64  `json:"burst-in-msec,omitempty"`
	CburstInMsec                  *int64  `json:"cburst-in-msec,omitempty"`
	ClassId                       *int64  `json:"class-id,omitempty"`
	GuaranteedBandwidthPercentage *int64  `json:"guaranteed-bandwidth-percentage,omitempty"`
	Id                            *int64  `json:"id,omitempty"`
	Limit                         *int64  `json:"limit,omitempty"`
	Max                           *int64  `json:"max,omitempty"`
	MaximumBandwidthPercentage    *int64  `json:"maximum-bandwidth-percentage,omitempty"`
	Min                           *int64  `json:"min,omitempty"`
	Priority                      *string `json:"priority,omitempty"`
	RedProbability                *int64  `json:"red-probability,omitempty"`
}
