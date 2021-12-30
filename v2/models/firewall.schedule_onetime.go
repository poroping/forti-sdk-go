package models

const FirewallscheduleOnetimePath = "firewall/schedule/onetime/"

type FirewallscheduleOnetime struct {
	Color          *int64  `json:"color,omitempty"`
	End            *string `json:"end,omitempty"`
	ExpirationDays *int64  `json:"expiration-days,omitempty"`
	FabricObject   *string `json:"fabric-object,omitempty"`
	Name           *string `json:"name,omitempty"`
	Start          *string `json:"start,omitempty"`
}
