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

const SystemAutomationStitchActionPath = "system/automation-stitch/action/"

type SystemAutomationStitchAction struct {
	Name *string `json:"name,omitempty"`
}

const SystemAutomationStitchActionsPath = "system/automation-stitch/actions/"

type SystemAutomationStitchActions struct {
	Action   *string `json:"action,omitempty"`
	Delay    *int64  `json:"delay,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	Required *string `json:"required,omitempty"`
}

const SystemAutomationStitchDestinationPath = "system/automation-stitch/destination/"

type SystemAutomationStitchDestination struct {
	Name *string `json:"name,omitempty"`
}
