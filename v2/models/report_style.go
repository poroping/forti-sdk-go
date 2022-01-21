package models

const ReportStylePath = "report/style/"

type ReportStyle struct {
	Align         *string `json:"align,omitempty"`
	BgColor       *string `json:"bg-color,omitempty"`
	BorderBottom  *string `json:"border-bottom,omitempty"`
	BorderLeft    *string `json:"border-left,omitempty"`
	BorderRight   *string `json:"border-right,omitempty"`
	BorderTop     *string `json:"border-top,omitempty"`
	ColumnGap     *string `json:"column-gap,omitempty"`
	ColumnSpan    *string `json:"column-span,omitempty"`
	FgColor       *string `json:"fg-color,omitempty"`
	FontFamily    *string `json:"font-family,omitempty"`
	FontSize      *string `json:"font-size,omitempty"`
	FontStyle     *string `json:"font-style,omitempty"`
	FontWeight    *string `json:"font-weight,omitempty"`
	Height        *string `json:"height,omitempty"`
	LineHeight    *string `json:"line-height,omitempty"`
	MarginBottom  *string `json:"margin-bottom,omitempty"`
	MarginLeft    *string `json:"margin-left,omitempty"`
	MarginRight   *string `json:"margin-right,omitempty"`
	MarginTop     *string `json:"margin-top,omitempty"`
	Name          *string `json:"name,omitempty"`
	Options       *string `json:"options,omitempty"`
	PaddingBottom *string `json:"padding-bottom,omitempty"`
	PaddingLeft   *string `json:"padding-left,omitempty"`
	PaddingRight  *string `json:"padding-right,omitempty"`
	PaddingTop    *string `json:"padding-top,omitempty"`
	Width         *string `json:"width,omitempty"`
}
