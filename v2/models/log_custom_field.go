package models

const LogCustomFieldPath = "log/custom-field/"

type LogCustomField struct {
	Fosid *string `json:"fosid,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
