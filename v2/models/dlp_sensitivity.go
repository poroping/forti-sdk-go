package models

const DlpSensitivityPath = "dlp/sensitivity/"

type DlpSensitivity struct {
	Name *string `json:"name,omitempty"`
}
