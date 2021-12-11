package models

const DlpSettingsPath = "dlp/settings/"

type DlpSettings struct {
	CacheMemPercent *int64  `json:"cache-mem-percent,omitempty"`
	ChunkSize       *int64  `json:"chunk-size,omitempty"`
	DbMode          *string `json:"db-mode,omitempty"`
	Size            *int64  `json:"size,omitempty"`
	StorageDevice   *string `json:"storage-device,omitempty"`
}
