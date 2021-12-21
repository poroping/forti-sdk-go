package models

const SystemAutoScriptPath = "system/auto-script/"

type SystemAutoScript struct {
	Interval   *float64 `json:"interval,omitempty"`
	Name       *string  `json:"name,omitempty"`
	OutputSize *float64 `json:"output-size,omitempty"`
	Repeat     *float64 `json:"repeat,omitempty"`
	Script     *string  `json:"script,omitempty"`
	Start      *string  `json:"start,omitempty"`
	Timeout    *float64 `json:"timeout,omitempty"`
}
