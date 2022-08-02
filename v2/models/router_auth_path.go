package models

const RouterAuthPathPath = "router/auth-path/"

type RouterAuthPath struct {
	Device  *string `json:"device,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Name    *string `json:"name,omitempty"`
}
