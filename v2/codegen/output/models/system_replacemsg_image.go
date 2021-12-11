package models

const SystemReplacemsgImagePath = "system/replacemsg-image/"

type SystemReplacemsgImage struct {
	ImageBase64 *string `json:"image-base64,omitempty"`
	ImageType   *string `json:"image-type,omitempty"`
	Name        *string `json:"name,omitempty"`
}
