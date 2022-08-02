package models

const VideofilterYoutubeKeyPath = "videofilter/youtube-key/"

type VideofilterYoutubeKey struct {
	Id  *int64  `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
}
