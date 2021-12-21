package models

const ReportChartPath = "report/chart/"

type ReportChart struct {
	Background      *string                       `json:"background,omitempty"`
	Category        *string                       `json:"category,omitempty"`
	CategorySeries  *[]ReportChartCategorySeries  `json:"category-series,omitempty"`
	ColorPalette    *string                       `json:"color-palette,omitempty"`
	Column          *[]ReportChartColumn          `json:"column,omitempty"`
	Comments        *string                       `json:"comments,omitempty"`
	Dataset         *string                       `json:"dataset,omitempty"`
	Dimension       *string                       `json:"dimension,omitempty"`
	DrillDownCharts *[]ReportChartDrillDownCharts `json:"drill-down-charts,omitempty"`
	Favorite        *string                       `json:"favorite,omitempty"`
	GraphType       *string                       `json:"graph-type,omitempty"`
	Legend          *string                       `json:"legend,omitempty"`
	LegendFontSize  *float64                      `json:"legend-font-size,omitempty"`
	Name            *string                       `json:"name,omitempty"`
	Period          *string                       `json:"period,omitempty"`
	Policy          *float64                      `json:"policy,omitempty"`
	Style           *string                       `json:"style,omitempty"`
	Title           *string                       `json:"title,omitempty"`
	TitleFontSize   *float64                      `json:"title-font-size,omitempty"`
	Type            *string                       `json:"type,omitempty"`
	ValueSeries     *[]ReportChartValueSeries     `json:"value-series,omitempty"`
	XSeries         *[]ReportChartXSeries         `json:"x-series,omitempty"`
	YSeries         *[]ReportChartYSeries         `json:"y-series,omitempty"`
}

type ReportChartCategorySeries struct {
	Databind *string  `json:"databind,omitempty"`
	FontSize *float64 `json:"font-size,omitempty"`
}

type ReportChartColumn struct {
	DetailUnit  *string                     `json:"detail-unit,omitempty"`
	DetailValue *string                     `json:"detail-value,omitempty"`
	FooterUnit  *string                     `json:"footer-unit,omitempty"`
	FooterValue *string                     `json:"footer-value,omitempty"`
	HeaderValue *string                     `json:"header-value,omitempty"`
	Id          *float64                    `json:"id,omitempty"`
	Mapping     *[]ReportChartColumnMapping `json:"mapping,omitempty"`
}

type ReportChartColumnMapping struct {
	Displayname *string  `json:"displayname,omitempty"`
	Id          *float64 `json:"id,omitempty"`
	Op          *string  `json:"op,omitempty"`
	ValueType   *string  `json:"value-type,omitempty"`
	Value1      *string  `json:"value1,omitempty"`
	Value2      *string  `json:"value2,omitempty"`
}

type ReportChartDrillDownCharts struct {
	ChartName *string  `json:"chart-name,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Status    *string  `json:"status,omitempty"`
}

type ReportChartValueSeries struct {
	Databind *string `json:"databind,omitempty"`
}

type ReportChartXSeries struct {
	Caption         *string  `json:"caption,omitempty"`
	CaptionFontSize *float64 `json:"caption-font-size,omitempty"`
	Databind        *string  `json:"databind,omitempty"`
	FontSize        *float64 `json:"font-size,omitempty"`
	IsCategory      *string  `json:"is-category,omitempty"`
	LabelAngle      *string  `json:"label-angle,omitempty"`
	ScaleDirection  *string  `json:"scale-direction,omitempty"`
	ScaleFormat     *string  `json:"scale-format,omitempty"`
	ScaleStep       *float64 `json:"scale-step,omitempty"`
	ScaleUnit       *string  `json:"scale-unit,omitempty"`
	Unit            *string  `json:"unit,omitempty"`
}

type ReportChartYSeries struct {
	Caption         *string  `json:"caption,omitempty"`
	CaptionFontSize *float64 `json:"caption-font-size,omitempty"`
	Databind        *string  `json:"databind,omitempty"`
	ExtraDatabind   *string  `json:"extra-databind,omitempty"`
	ExtraY          *string  `json:"extra-y,omitempty"`
	ExtraYLegend    *string  `json:"extra-y-legend,omitempty"`
	FontSize        *float64 `json:"font-size,omitempty"`
	Group           *string  `json:"group,omitempty"`
	LabelAngle      *string  `json:"label-angle,omitempty"`
	Unit            *string  `json:"unit,omitempty"`
	YLegend         *string  `json:"y-legend,omitempty"`
}
