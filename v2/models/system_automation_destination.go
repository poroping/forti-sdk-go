package models

const SystemAutomationDestinationPath = "system/automation-destination/"

type SystemAutomationDestination struct {
	Destination *[]SystemAutomationDestinationDestination `json:"destination,omitempty"`
	HaGroupId   *float64                                  `json:"ha-group-id,omitempty"`
	Name        *string                                   `json:"name,omitempty"`
	Type        *string                                   `json:"type,omitempty"`
}

type SystemAutomationDestinationDestination struct {
	Name *string `json:"name,omitempty"`
}
