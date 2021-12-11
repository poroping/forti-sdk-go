package models

const VideofilterYoutubeKeyPath = "videofilter/youtube-key/"

type VideofilterYoutubeKey struct {
	Fosid *int64  `json:"fosid,omitempty"`
	Key   *string `json:"key,omitempty"`
}
