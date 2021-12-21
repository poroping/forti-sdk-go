package models

const WebfilterFtgdLocalCatPath = "webfilter/ftgd-local-cat/"

type WebfilterFtgdLocalCat struct {
	Desc   *string  `json:"desc,omitempty"`
	Fosid  *float64 `json:"fosid,omitempty"`
	Status *string  `json:"status,omitempty"`
}
