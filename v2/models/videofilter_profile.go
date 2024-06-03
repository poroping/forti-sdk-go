package models

const VideofilterProfilePath = "videofilter/profile/"

type VideofilterProfile struct {
	Comment              *string                               `json:"comment,omitempty"`
	Dailymotion          *string                               `json:"dailymotion,omitempty"`
	DefaultAction        *string                               `json:"default-action,omitempty"`
	FortiguardCategory   *VideofilterProfileFortiguardCategory `json:"fortiguard-category,omitempty"`
	Log                  *string                               `json:"log,omitempty"`
	Name                 *string                               `json:"name,omitempty"`
	ReplacemsgGroup      *string                               `json:"replacemsg-group,omitempty"`
	Vimeo                *string                               `json:"vimeo,omitempty"`
	VimeoRestrict        *string                               `json:"vimeo-restrict,omitempty"`
	Youtube              *string                               `json:"youtube,omitempty"`
	YoutubeChannelFilter *int64                                `json:"youtube-channel-filter,omitempty"`
	YoutubeRestrict      *string                               `json:"youtube-restrict,omitempty"`
}

type VideofilterProfileFortiguardCategory struct {
	Filters *[]VideofilterProfileFortiguardCategoryFilters `json:"filters,omitempty"`
}

type VideofilterProfileFortiguardCategoryFilters struct {
	Action     *string `json:"action,omitempty"`
	CategoryId *int64  `json:"category-id,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Log        *string `json:"log,omitempty"`
}
