package models

const IpsDecoderPath = "ips/decoder/"

type IpsDecoder struct {
	Name      *string                `json:"name,omitempty"`
	Parameter *[]IpsDecoderParameter `json:"parameter,omitempty"`
}

type IpsDecoderParameter struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
