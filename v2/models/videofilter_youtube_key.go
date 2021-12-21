package models

const VideofilterYoutubeKeyPath = "videofilter/youtube-key/"

type VideofilterYoutubeKey struct {
	Fosid  *float64 `json:"fosid,omitempty"`
	Key    *string  `json:"key,omitempty"`
	Status *string  `json:"status,omitempty"`
}
