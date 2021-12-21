package models

const VideofilterProfilePath = "videofilter/profile/"

type VideofilterProfile struct {
	Comment              *string                                 `json:"comment,omitempty"`
	Dailymotion          *string                                 `json:"dailymotion,omitempty"`
	FortiguardCategory   *[]VideofilterProfileFortiguardCategory `json:"fortiguard-category,omitempty"`
	Name                 *string                                 `json:"name,omitempty"`
	ReplacemsgGroup      *string                                 `json:"replacemsg-group,omitempty"`
	Vimeo                *string                                 `json:"vimeo,omitempty"`
	VimeoRestrict        *string                                 `json:"vimeo-restrict,omitempty"`
	Youtube              *string                                 `json:"youtube,omitempty"`
	YoutubeChannelFilter *float64                                `json:"youtube-channel-filter,omitempty"`
	YoutubeRestrict      *string                                 `json:"youtube-restrict,omitempty"`
}

type VideofilterProfileFortiguardCategory struct {
	Filters *[]VideofilterProfileFortiguardCategoryFilters `json:"filters,omitempty"`
}

type VideofilterProfileFortiguardCategoryFilters struct {
	Action     *string  `json:"action,omitempty"`
	CategoryId *float64 `json:"category-id,omitempty"`
	Id         *float64 `json:"id,omitempty"`
	Log        *string  `json:"log,omitempty"`
}
