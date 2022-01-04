package models

const FirewallScheduleRecurringPath = "firewall.schedule/recurring/"

type FirewallScheduleRecurring struct {
	Color        *int64  `json:"color,omitempty"`
	Day          *string `json:"day,omitempty"`
	End          *string `json:"end,omitempty"`
	FabricObject *string `json:"fabric-object,omitempty"`
	Name         *string `json:"name,omitempty"`
	Start        *string `json:"start,omitempty"`
}
