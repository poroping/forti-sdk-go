package models

const SystemAutoScriptPath = "system/auto-script/"

type SystemAutoScript struct {
	Interval   *int64  `json:"interval,omitempty"`
	Name       *string `json:"name,omitempty"`
	OutputSize *int64  `json:"output-size,omitempty"`
	Repeat     *int64  `json:"repeat,omitempty"`
	Script     *string `json:"script,omitempty"`
	Start      *string `json:"start,omitempty"`
	Timeout    *int64  `json:"timeout,omitempty"`
}
