package models

const SystemStoragePath = "system/storage/"

type SystemStorage struct {
	Device      *string  `json:"device,omitempty"`
	MediaStatus *string  `json:"media-status,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Order       *float64 `json:"order,omitempty"`
	Partition   *string  `json:"partition,omitempty"`
	Size        *float64 `json:"size,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Usage       *string  `json:"usage,omitempty"`
	WanoptMode  *string  `json:"wanopt-mode,omitempty"`
}
