package models

const WebfilterFtgdLocalRatingPath = "webfilter/ftgd-local-rating/"

type WebfilterFtgdLocalRating struct {
	Comment *string `json:"comment,omitempty"`
	Rating  *string `json:"rating,omitempty"`
	Status  *string `json:"status,omitempty"`
	Url     *string `json:"url,omitempty"`
}
