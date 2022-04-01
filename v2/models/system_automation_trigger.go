package models

const SystemAutomationTriggerPath = "system/automation-trigger/"

type SystemAutomationTrigger struct {
	Description         *string                          `json:"description,omitempty"`
	EventType           *string                          `json:"event-type,omitempty"`
	FabricEventName     *string                          `json:"fabric-event-name,omitempty"`
	FabricEventSeverity *string                          `json:"fabric-event-severity,omitempty"`
	FazEventName        *string                          `json:"faz-event-name,omitempty"`
	FazEventSeverity    *string                          `json:"faz-event-severity,omitempty"`
	FazEventTags        *string                          `json:"faz-event-tags,omitempty"`
	Fields              *[]SystemAutomationTriggerFields `json:"fields,omitempty"`
	IocLevel            *string                          `json:"ioc-level,omitempty"`
	LicenseType         *string                          `json:"license-type,omitempty"`
	Logid               *[]SystemAutomationTriggerLogid  `json:"logid,omitempty"`
	Name                *string                          `json:"name,omitempty"`
	ReportType          *string                          `json:"report-type,omitempty"`
	Serial              *string                          `json:"serial,omitempty"`
	TriggerDay          *int64                           `json:"trigger-day,omitempty"`
	TriggerFrequency    *string                          `json:"trigger-frequency,omitempty"`
	TriggerHour         *int64                           `json:"trigger-hour,omitempty"`
	TriggerMinute       *int64                           `json:"trigger-minute,omitempty"`
	TriggerType         *string                          `json:"trigger-type,omitempty"`
	TriggerWeekday      *string                          `json:"trigger-weekday,omitempty"`
	Vdom                *[]SystemAutomationTriggerVdom   `json:"vdom,omitempty"`
}

type SystemAutomationTriggerFields struct {
	Id    *int64  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type SystemAutomationTriggerLogid struct {
	Id *int64 `json:"id,omitempty"`
}

type SystemAutomationTriggerVdom struct {
	Name *string `json:"name,omitempty"`
}
