package models

const SystemIkePath = "system/ike/"

type SystemIke struct {
	DhGroup1          *[]SystemIkeDhGroup1  `json:"dh-group-1,omitempty"`
	DhGroup14         *[]SystemIkeDhGroup14 `json:"dh-group-14,omitempty"`
	DhGroup15         *[]SystemIkeDhGroup15 `json:"dh-group-15,omitempty"`
	DhGroup16         *[]SystemIkeDhGroup16 `json:"dh-group-16,omitempty"`
	DhGroup17         *[]SystemIkeDhGroup17 `json:"dh-group-17,omitempty"`
	DhGroup18         *[]SystemIkeDhGroup18 `json:"dh-group-18,omitempty"`
	DhGroup19         *[]SystemIkeDhGroup19 `json:"dh-group-19,omitempty"`
	DhGroup2          *[]SystemIkeDhGroup2  `json:"dh-group-2,omitempty"`
	DhGroup20         *[]SystemIkeDhGroup20 `json:"dh-group-20,omitempty"`
	DhGroup21         *[]SystemIkeDhGroup21 `json:"dh-group-21,omitempty"`
	DhGroup27         *[]SystemIkeDhGroup27 `json:"dh-group-27,omitempty"`
	DhGroup28         *[]SystemIkeDhGroup28 `json:"dh-group-28,omitempty"`
	DhGroup29         *[]SystemIkeDhGroup29 `json:"dh-group-29,omitempty"`
	DhGroup30         *[]SystemIkeDhGroup30 `json:"dh-group-30,omitempty"`
	DhGroup31         *[]SystemIkeDhGroup31 `json:"dh-group-31,omitempty"`
	DhGroup32         *[]SystemIkeDhGroup32 `json:"dh-group-32,omitempty"`
	DhGroup5          *[]SystemIkeDhGroup5  `json:"dh-group-5,omitempty"`
	DhKeypairCache    *string               `json:"dh-keypair-cache,omitempty"`
	DhKeypairCount    *int64                `json:"dh-keypair-count,omitempty"`
	DhKeypairThrottle *string               `json:"dh-keypair-throttle,omitempty"`
	DhMode            *string               `json:"dh-mode,omitempty"`
	DhMultiprocess    *string               `json:"dh-multiprocess,omitempty"`
	DhWorkerCount     *int64                `json:"dh-worker-count,omitempty"`
	EmbryonicLimit    *int64                `json:"embryonic-limit,omitempty"`
}

type SystemIkeDhGroup1 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup14 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup15 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup16 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup17 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup18 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup19 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup2 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup20 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup21 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup27 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup28 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup29 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup30 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup31 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup32 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

type SystemIkeDhGroup5 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}
