package models

const SystemAliasPath = "system/alias/"

type SystemAlias struct {
	Command *string `json:"command,omitempty"`
	Name    *string `json:"name,omitempty"`
}
