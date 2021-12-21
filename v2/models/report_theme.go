package models

const ReportThemePath = "report/theme/"

type ReportTheme struct {
	BulletListStyle        *string `json:"bullet-list-style,omitempty"`
	ColumnCount            *string `json:"column-count,omitempty"`
	DefaultHtmlStyle       *string `json:"default-html-style,omitempty"`
	DefaultPdfStyle        *string `json:"default-pdf-style,omitempty"`
	GraphChartStyle        *string `json:"graph-chart-style,omitempty"`
	Heading1Style          *string `json:"heading1-style,omitempty"`
	Heading2Style          *string `json:"heading2-style,omitempty"`
	Heading3Style          *string `json:"heading3-style,omitempty"`
	Heading4Style          *string `json:"heading4-style,omitempty"`
	HlineStyle             *string `json:"hline-style,omitempty"`
	ImageStyle             *string `json:"image-style,omitempty"`
	Name                   *string `json:"name,omitempty"`
	NormalTextStyle        *string `json:"normal-text-style,omitempty"`
	NumberedListStyle      *string `json:"numbered-list-style,omitempty"`
	PageFooterStyle        *string `json:"page-footer-style,omitempty"`
	PageHeaderStyle        *string `json:"page-header-style,omitempty"`
	PageOrient             *string `json:"page-orient,omitempty"`
	PageStyle              *string `json:"page-style,omitempty"`
	ReportSubtitleStyle    *string `json:"report-subtitle-style,omitempty"`
	ReportTitleStyle       *string `json:"report-title-style,omitempty"`
	TableChartCaptionStyle *string `json:"table-chart-caption-style,omitempty"`
	TableChartEvenRowStyle *string `json:"table-chart-even-row-style,omitempty"`
	TableChartHeadStyle    *string `json:"table-chart-head-style,omitempty"`
	TableChartOddRowStyle  *string `json:"table-chart-odd-row-style,omitempty"`
	TableChartStyle        *string `json:"table-chart-style,omitempty"`
	TocHeading1Style       *string `json:"toc-heading1-style,omitempty"`
	TocHeading2Style       *string `json:"toc-heading2-style,omitempty"`
	TocHeading3Style       *string `json:"toc-heading3-style,omitempty"`
	TocHeading4Style       *string `json:"toc-heading4-style,omitempty"`
	TocTitleStyle          *string `json:"toc-title-style,omitempty"`
}
