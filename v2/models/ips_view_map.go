package models

const IpsViewMapPath = "ips/view-map/"

type IpsViewMap struct {
	Id         *int64  `json:"id,omitempty"`
	IdPolicyId *int64  `json:"id-policy-id,omitempty"`
	PolicyId   *int64  `json:"policy-id,omitempty"`
	VdomId     *int64  `json:"vdom-id,omitempty"`
	Which      *string `json:"which,omitempty"`
}
