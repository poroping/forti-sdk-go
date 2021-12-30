package models

const WebfilterFtgdLocalCatPath = "webfilter/ftgd-local-cat/"

type WebfilterFtgdLocalCat struct {
	Desc   *string `json:"desc,omitempty"`
	Fosid  *int64  `json:"fosid,omitempty"`
	Status *string `json:"status,omitempty"`
}
