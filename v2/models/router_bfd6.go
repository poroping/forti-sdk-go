package models

const RouterBfd6Path = "router/bfd6/"

type RouterBfd6 struct {
	MultihopTemplate *[]RouterBfd6MultihopTemplate `json:"multihop-template,omitempty"`
	Neighbor         *[]RouterBfd6Neighbor         `json:"neighbor,omitempty"`
}

type RouterBfd6MultihopTemplate struct {
	AuthMode         *string `json:"auth-mode,omitempty"`
	BfdDesiredMinTx  *int64  `json:"bfd-desired-min-tx,omitempty"`
	BfdDetectMult    *int64  `json:"bfd-detect-mult,omitempty"`
	BfdRequiredMinRx *int64  `json:"bfd-required-min-rx,omitempty"`
	Dst              *string `json:"dst,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	Md5Key           *string `json:"md5-key,omitempty"`
	Src              *string `json:"src,omitempty"`
}

type RouterBfd6Neighbor struct {
	Interface  *string `json:"interface,omitempty"`
	Ip6Address *string `json:"ip6-address,omitempty"`
}
