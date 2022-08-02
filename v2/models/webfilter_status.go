package models

const WebfilterStatusPath = "webfilter/status/"

type WebfilterStatus struct {
	RefreshRate *string `json:"<refresh-rate>,omitempty"`
}
