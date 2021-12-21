package models

const IpsViewMapPath = "ips/view-map/"

type IpsViewMap struct {
	Fosid      *float64 `json:"fosid,omitempty"`
	IdPolicyId *float64 `json:"id-policy-id,omitempty"`
	PolicyId   *float64 `json:"policy-id,omitempty"`
	VdomId     *float64 `json:"vdom-id,omitempty"`
	Which      *string  `json:"which,omitempty"`
}
