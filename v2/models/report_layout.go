package models

const ReportLayoutPath = "report/layout/"

type ReportLayout struct {
	BodyItem        *[]ReportLayoutBodyItem `json:"body-item,omitempty"`
	CutoffOption    *string                 `json:"cutoff-option,omitempty"`
	CutoffTime      *string                 `json:"cutoff-time,omitempty"`
	Day             *string                 `json:"day,omitempty"`
	Description     *string                 `json:"description,omitempty"`
	EmailRecipients *string                 `json:"email-recipients,omitempty"`
	EmailSend       *string                 `json:"email-send,omitempty"`
	Format          *string                 `json:"format,omitempty"`
	MaxPdfReport    *int64                  `json:"max-pdf-report,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Options         *string                 `json:"options,omitempty"`
	Page            *[]ReportLayoutPage     `json:"page,omitempty"`
	ScheduleType    *string                 `json:"schedule-type,omitempty"`
	StyleTheme      *string                 `json:"style-theme,omitempty"`
	Subtitle        *string                 `json:"subtitle,omitempty"`
	Time            *string                 `json:"time,omitempty"`
	Title           *string                 `json:"title,omitempty"`
}

type ReportLayoutBodyItem struct {
	Chart             *string                           `json:"chart,omitempty"`
	ChartOptions      *string                           `json:"chart-options,omitempty"`
	Column            *int64                            `json:"column,omitempty"`
	Content           *string                           `json:"content,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	DrillDownItems    *string                           `json:"drill-down-items,omitempty"`
	DrillDownTypes    *string                           `json:"drill-down-types,omitempty"`
	Hide              *string                           `json:"hide,omitempty"`
	Id                *int64                            `json:"id,omitempty"`
	ImgSrc            *string                           `json:"img-src,omitempty"`
	List              *[]ReportLayoutBodyItemList       `json:"list,omitempty"`
	ListComponent     *string                           `json:"list-component,omitempty"`
	MiscComponent     *string                           `json:"misc-component,omitempty"`
	Parameters        *[]ReportLayoutBodyItemParameters `json:"parameters,omitempty"`
	Style             *string                           `json:"style,omitempty"`
	TableCaptionStyle *string                           `json:"table-caption-style,omitempty"`
	TableColumnWidths *string                           `json:"table-column-widths,omitempty"`
	TableEvenRowStyle *string                           `json:"table-even-row-style,omitempty"`
	TableHeadStyle    *string                           `json:"table-head-style,omitempty"`
	TableOddRowStyle  *string                           `json:"table-odd-row-style,omitempty"`
	TextComponent     *string                           `json:"text-component,omitempty"`
	Title             *string                           `json:"title,omitempty"`
	TopN              *int64                            `json:"top-n,omitempty"`
	Type              *string                           `json:"type,omitempty"`
}

type ReportLayoutBodyItemList struct {
	Content *string `json:"content,omitempty"`
	Id      *int64  `json:"id,omitempty"`
}

type ReportLayoutBodyItemParameters struct {
	Id    *int64  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ReportLayoutPage struct {
	ColumnBreakBefore *string                   `json:"column-break-before,omitempty"`
	Footer            *[]ReportLayoutPageFooter `json:"footer,omitempty"`
	Header            *[]ReportLayoutPageHeader `json:"header,omitempty"`
	Options           *string                   `json:"options,omitempty"`
	PageBreakBefore   *string                   `json:"page-break-before,omitempty"`
	Paper             *string                   `json:"paper,omitempty"`
}

type ReportLayoutPageFooter struct {
	FooterItem *[]ReportLayoutPageFooterFooterItem `json:"footer-item,omitempty"`
	Style      *string                             `json:"style,omitempty"`
}

type ReportLayoutPageFooterFooterItem struct {
	Content     *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	ImgSrc      *string `json:"img-src,omitempty"`
	Style       *string `json:"style,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type ReportLayoutPageHeader struct {
	HeaderItem *[]ReportLayoutPageHeaderHeaderItem `json:"header-item,omitempty"`
	Style      *string                             `json:"style,omitempty"`
}

type ReportLayoutPageHeaderHeaderItem struct {
	Content     *string `json:"content,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	ImgSrc      *string `json:"img-src,omitempty"`
	Style       *string `json:"style,omitempty"`
	Type        *string `json:"type,omitempty"`
}
