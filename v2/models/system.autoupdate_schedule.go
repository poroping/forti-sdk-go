package models

const SystemautoupdateSchedulePath = "system.autoupdate/schedule/"

type SystemautoupdateSchedule struct {
	Day       *string `json:"day,omitempty"`
	Frequency *string `json:"frequency,omitempty"`
	Status    *string `json:"status,omitempty"`
	Time      *string `json:"time,omitempty"`
}
