package models

const ApplicationCustomPath = "application/custom/"

type ApplicationCustom struct {
	Behavior   *string `json:"behavior,omitempty"`
	Category   *int64  `json:"category,omitempty"`
	Comment    *string `json:"comment,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Protocol   *string `json:"protocol,omitempty"`
	Signature  *string `json:"signature,omitempty"`
	Tag        *string `json:"tag,omitempty"`
	Technology *string `json:"technology,omitempty"`
	Vendor     *string `json:"vendor,omitempty"`
}
