package models

const DlpSensorPath = "dlp/sensor/"

type DlpSensor struct {
	Comment          *string            `json:"comment,omitempty"`
	DlpLog           *string            `json:"dlp-log,omitempty"`
	ExtendedLog      *string            `json:"extended-log,omitempty"`
	FeatureSet       *string            `json:"feature-set,omitempty"`
	Filter           *[]DlpSensorFilter `json:"filter,omitempty"`
	FullArchiveProto *string            `json:"full-archive-proto,omitempty"`
	NacQuarLog       *string            `json:"nac-quar-log,omitempty"`
	Name             *string            `json:"name,omitempty"`
	Options          *string            `json:"options,omitempty"`
	ReplacemsgGroup  *string            `json:"replacemsg-group,omitempty"`
	SummaryProto     *string            `json:"summary-proto,omitempty"`
}

type DlpSensorFilter struct {
	Action            *string                       `json:"action,omitempty"`
	Archive           *string                       `json:"archive,omitempty"`
	CompanyIdentifier *string                       `json:"company-identifier,omitempty"`
	Expiry            *string                       `json:"expiry,omitempty"`
	FileSize          *float64                      `json:"file-size,omitempty"`
	FileType          *float64                      `json:"file-type,omitempty"`
	FilterBy          *string                       `json:"filter-by,omitempty"`
	Id                *float64                      `json:"id,omitempty"`
	MatchPercentage   *float64                      `json:"match-percentage,omitempty"`
	Name              *string                       `json:"name,omitempty"`
	Proto             *string                       `json:"proto,omitempty"`
	Regexp            *string                       `json:"regexp,omitempty"`
	Sensitivity       *[]DlpSensorFilterSensitivity `json:"sensitivity,omitempty"`
	Severity          *string                       `json:"severity,omitempty"`
	Type              *string                       `json:"type,omitempty"`
}

type DlpSensorFilterSensitivity struct {
	Name *string `json:"name,omitempty"`
}
