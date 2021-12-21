package models

const SystemsdwanMembersPath = "system.sdwan/members/"

type SystemsdwanMembers struct {
	Comment                   *string `json:"comment,omitempty"`
	Cost                      *int64  `json:"cost,omitempty"`
	Gateway                   *string `json:"gateway,omitempty"`
	Gateway6                  *string `json:"gateway6,omitempty"`
	IngressSpilloverThreshold *int64  `json:"ingress-spillover-threshold,omitempty"`
	Interface                 *string `json:"interface,omitempty"`
	Priority                  *int64  `json:"priority,omitempty"`
	Priority6                 *int64  `json:"priority6,omitempty"`
	SeqNum                    *int64  `json:"seq-num,omitempty"`
	Source                    *string `json:"source,omitempty"`
	Source6                   *string `json:"source6,omitempty"`
	SpilloverThreshold        *int64  `json:"spillover-threshold,omitempty"`
	Status                    *string `json:"status,omitempty"`
	VolumeRatio               *int64  `json:"volume-ratio,omitempty"`
	Weight                    *int64  `json:"weight,omitempty"`
	Zone                      *string `json:"zone,omitempty"`
}
