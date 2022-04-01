package models

const RouterBfdPath = "router/bfd/"

type RouterBfd struct {
	MultihopTemplate *[]RouterBfdMultihopTemplate `json:"multihop-template,omitempty"`
	Neighbor         *[]RouterBfdNeighbor         `json:"neighbor,omitempty"`
}

type RouterBfdMultihopTemplate struct {
	AuthMode         *string `json:"auth-mode,omitempty"`
	BfdDesiredMinTx  *int64  `json:"bfd-desired-min-tx,omitempty"`
	BfdDetectMult    *int64  `json:"bfd-detect-mult,omitempty"`
	BfdRequiredMinRx *int64  `json:"bfd-required-min-rx,omitempty"`
	Dst              *string `json:"dst,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	Md5Key           *string `json:"md5-key,omitempty"`
	Src              *string `json:"src,omitempty"`
}

type RouterBfdNeighbor struct {
	Interface *string `json:"interface,omitempty"`
	Ip        *string `json:"ip,omitempty"`
}
