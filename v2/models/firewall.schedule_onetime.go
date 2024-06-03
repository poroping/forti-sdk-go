package models

const FirewallScheduleOnetimePath = "firewall.schedule/onetime/"

type FirewallScheduleOnetime struct {
	Color          *int64  `json:"color,omitempty"`
	End            *string `json:"end,omitempty"`
	EndUtc         *string `json:"end-utc,omitempty"`
	ExpirationDays *int64  `json:"expiration-days,omitempty"`
	FabricObject   *string `json:"fabric-object,omitempty"`
	Name           *string `json:"name,omitempty"`
	Start          *string `json:"start,omitempty"`
	StartUtc       *string `json:"start-utc,omitempty"`
}
