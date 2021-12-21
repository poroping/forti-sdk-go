package models

const SystemAutomationStitchPath = "system/automation-stitch/"

type SystemAutomationStitch struct {
	Action      *[]SystemAutomationStitchAction      `json:"action,omitempty"`
	Actions     *[]SystemAutomationStitchActions     `json:"actions,omitempty"`
	Description *string                              `json:"description,omitempty"`
	Destination *[]SystemAutomationStitchDestination `json:"destination,omitempty"`
	Name        *string                              `json:"name,omitempty"`
	Status      *string                              `json:"status,omitempty"`
	Trigger     *string                              `json:"trigger,omitempty"`
}

type SystemAutomationStitchAction struct {
	Name *string `json:"name,omitempty"`
}

type SystemAutomationStitchActions struct {
	Action   *string  `json:"action,omitempty"`
	Delay    *float64 `json:"delay,omitempty"`
	Id       *float64 `json:"id,omitempty"`
	Required *string  `json:"required,omitempty"`
}

type SystemAutomationStitchDestination struct {
	Name *string `json:"name,omitempty"`
}
