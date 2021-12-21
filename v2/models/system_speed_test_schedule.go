package models

const SystemSpeedTestSchedulePath = "system/speed-test-schedule/"

type SystemSpeedTestSchedule struct {
	Diffserv                  *string                             `json:"diffserv,omitempty"`
	DynamicServer             *string                             `json:"dynamic-server,omitempty"`
	Interface                 *string                             `json:"interface,omitempty"`
	Schedules                 *[]SystemSpeedTestScheduleSchedules `json:"schedules,omitempty"`
	ServerName                *string                             `json:"server-name,omitempty"`
	Status                    *string                             `json:"status,omitempty"`
	UpdateInbandwidth         *string                             `json:"update-inbandwidth,omitempty"`
	UpdateInbandwidthMaximum  *float64                            `json:"update-inbandwidth-maximum,omitempty"`
	UpdateInbandwidthMinimum  *float64                            `json:"update-inbandwidth-minimum,omitempty"`
	UpdateOutbandwidth        *string                             `json:"update-outbandwidth,omitempty"`
	UpdateOutbandwidthMaximum *float64                            `json:"update-outbandwidth-maximum,omitempty"`
	UpdateOutbandwidthMinimum *float64                            `json:"update-outbandwidth-minimum,omitempty"`
}

type SystemSpeedTestScheduleSchedules struct {
	Name *string `json:"name,omitempty"`
}
