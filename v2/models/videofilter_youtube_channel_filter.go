package models

const VideofilterYoutubeChannelFilterPath = "videofilter/youtube-channel-filter/"

type VideofilterYoutubeChannelFilter struct {
	Comment       *string                                   `json:"comment,omitempty"`
	DefaultAction *string                                   `json:"default-action,omitempty"`
	Entries       *[]VideofilterYoutubeChannelFilterEntries `json:"entries,omitempty"`
	Fosid         *float64                                  `json:"fosid,omitempty"`
	Log           *string                                   `json:"log,omitempty"`
	Name          *string                                   `json:"name,omitempty"`
}

type VideofilterYoutubeChannelFilterEntries struct {
	Action    *string  `json:"action,omitempty"`
	ChannelId *string  `json:"channel-id,omitempty"`
	Comment   *string  `json:"comment,omitempty"`
	Id        *float64 `json:"id,omitempty"`
}
