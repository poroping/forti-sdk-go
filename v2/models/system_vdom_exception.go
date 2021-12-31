package models

const SystemVdomExceptionPath = "system/vdom-exception/"

type SystemVdomException struct {
	Id     *int64                     `json:"id,omitempty"`
	Object *string                    `json:"object,omitempty"`
	Scope  *string                    `json:"scope,omitempty"`
	Vdom   *[]SystemVdomExceptionVdom `json:"vdom,omitempty"`
}

type SystemVdomExceptionVdom struct {
	Name *string `json:"name,omitempty"`
}
