package models

const DlpSettingsPath = "dlp/settings/"

type DlpSettings struct {
	CacheMemPercent *float64 `json:"cache-mem-percent,omitempty"`
	ChunkSize       *float64 `json:"chunk-size,omitempty"`
	DbMode          *string  `json:"db-mode,omitempty"`
	Size            *float64 `json:"size,omitempty"`
	StorageDevice   *string  `json:"storage-device,omitempty"`
}
