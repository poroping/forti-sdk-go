package models

const IpsDecoderPath = "ips/decoder/"

type IpsDecoder struct {
	Name      *string                `json:"name,omitempty"`
	Parameter *[]IpsDecoderParameter `json:"parameter,omitempty"`
}

const IpsDecoderParameterPath = "ips/decoder/parameter/"

type IpsDecoderParameter struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Set IpsDecoderParameter values to defaults
func (def *IpsDecoderParameter) Defaults() {
	def.Name = stringPtr("")
	def.Value = stringPtr("")
}

// Set IpsDecoder values to defaults
func (def *IpsDecoder) Defaults() {
	def.Name = stringPtr("")
}
