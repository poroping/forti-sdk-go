package models

const LogCustomFieldPath = "log/custom-field/"

type LogCustomField struct {
	Id    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
