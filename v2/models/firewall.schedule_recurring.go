package models

const FirewallscheduleRecurringPath = "firewall.schedule/recurring/"

type FirewallscheduleRecurring struct {
	Color        *float64 `json:"color,omitempty"`
	Day          *string  `json:"day,omitempty"`
	End          *string  `json:"end,omitempty"`
	FabricObject *string  `json:"fabric-object,omitempty"`
	Name         *string  `json:"name,omitempty"`
	Start        *string  `json:"start,omitempty"`
}
