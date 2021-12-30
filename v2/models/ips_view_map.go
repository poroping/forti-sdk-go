package models

const IpsViewMapPath = "ips/view-map/"

type IpsViewMap struct {
	Fosid      *int64  `json:"fosid,omitempty"`
	IdPolicyId *int64  `json:"id-policy-id,omitempty"`
	PolicyId   *int64  `json:"policy-id,omitempty"`
	VdomId     *int64  `json:"vdom-id,omitempty"`
	Which      *string `json:"which,omitempty"`
}
