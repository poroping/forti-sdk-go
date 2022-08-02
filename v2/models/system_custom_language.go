package models

const SystemCustomLanguagePath = "system/custom-language/"

type SystemCustomLanguage struct {
	Comments *string `json:"comments,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Name     *string `json:"name,omitempty"`
}
