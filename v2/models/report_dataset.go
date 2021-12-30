package models

const ReportDatasetPath = "report/dataset/"

type ReportDataset struct {
	Field      *[]ReportDatasetField      `json:"field,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Parameters *[]ReportDatasetParameters `json:"parameters,omitempty"`
	Policy     *int64                     `json:"policy,omitempty"`
	Query      *string                    `json:"query,omitempty"`
}

type ReportDatasetField struct {
	Displayname *string `json:"displayname,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type ReportDatasetParameters struct {
	DataType    *string `json:"data-type,omitempty"`
	DisplayName *string `json:"display-name,omitempty"`
	Field       *string `json:"field,omitempty"`
	Id          *int64  `json:"id,omitempty"`
}
