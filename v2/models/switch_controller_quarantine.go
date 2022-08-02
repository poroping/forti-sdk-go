package models

const SwitchControllerQuarantinePath = "switch-controller/quarantine/"

type SwitchControllerQuarantine struct {
	Quarantine *string                              `json:"quarantine,omitempty"`
	Targets    *[]SwitchControllerQuarantineTargets `json:"targets,omitempty"`
}

const SwitchControllerQuarantineTargetsPath = "switch-controller/quarantine/targets/"

type SwitchControllerQuarantineTargets struct {
	Description *string                                 `json:"description,omitempty"`
	Mac         *string                                 `json:"mac,omitempty"`
	Tag         *[]SwitchControllerQuarantineTargetsTag `json:"tag,omitempty"`
}

const SwitchControllerQuarantineTargetsTagPath = "switch-controller/quarantine/targets/tag/"

type SwitchControllerQuarantineTargetsTag struct {
	Tags *string `json:"tags,omitempty"`
}
