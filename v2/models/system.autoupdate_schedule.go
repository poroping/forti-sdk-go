package models

const SystemAutoupdateSchedulePath = "system.autoupdate/schedule/"

type SystemAutoupdateSchedule struct {
	Day       *string `json:"day,omitempty"`
	Frequency *string `json:"frequency,omitempty"`
	Status    *string `json:"status,omitempty"`
	Time      *string `json:"time,omitempty"`
}
