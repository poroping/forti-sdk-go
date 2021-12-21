package models

const AntivirusSettingsPath = "antivirus/settings/"

type AntivirusSettings struct {
	CacheCleanResult         *string `json:"cache-clean-result,omitempty"`
	CacheInfectedResult      *string `json:"cache-infected-result,omitempty"`
	DefaultDb                *string `json:"default-db,omitempty"`
	Grayware                 *string `json:"grayware,omitempty"`
	MachineLearningDetection *string `json:"machine-learning-detection,omitempty"`
	OverrideTimeout          *int64  `json:"override-timeout,omitempty"`
	UseExtremeDb             *string `json:"use-extreme-db,omitempty"`
}
