package models

const VideofilterProfilePath = "videofilter/profile/"

type VideofilterProfile struct {
	Comment              *string                               `json:"comment,omitempty"`
	Dailymotion          *string                               `json:"dailymotion,omitempty"`
	FortiguardCategory   *VideofilterProfileFortiguardCategory `json:"fortiguard-category,omitempty"`
	Name                 *string                               `json:"name,omitempty"`
	ReplacemsgGroup      *string                               `json:"replacemsg-group,omitempty"`
	Vimeo                *string                               `json:"vimeo,omitempty"`
	Youtube              *string                               `json:"youtube,omitempty"`
	YoutubeChannelFilter *int64                                `json:"youtube-channel-filter,omitempty"`
}

const VideofilterProfileFortiguardCategoryPath = "videofilter/profile/fortiguard-category/"

type VideofilterProfileFortiguardCategory struct {
	Filters *[]VideofilterProfileFortiguardCategoryFilters `json:"filters,omitempty"`
}

const VideofilterProfileFortiguardCategoryFiltersPath = "videofilter/profile/fortiguard-category/filters/"

type VideofilterProfileFortiguardCategoryFilters struct {
	Action     *string `json:"action,omitempty"`
	CategoryId *int64  `json:"category-id,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	Log        *string `json:"log,omitempty"`
}
