package models

const SystemVdomExceptionPath = "system/vdom-exception/"

type SystemVdomException struct {
	Fosid  *float64                   `json:"fosid,omitempty"`
	Object *string                    `json:"object,omitempty"`
	Scope  *string                    `json:"scope,omitempty"`
	Vdom   *[]SystemVdomExceptionVdom `json:"vdom,omitempty"`
}

type SystemVdomExceptionVdom struct {
	Name *string `json:"name,omitempty"`
}
