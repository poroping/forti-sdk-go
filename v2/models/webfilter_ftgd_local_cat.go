package models

const WebfilterFtgdLocalCatPath = "webfilter/ftgd-local-cat/"

type WebfilterFtgdLocalCat struct {
	Desc   *string `json:"desc,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}
