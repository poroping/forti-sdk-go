package models

const SystemIkePath = "system/ike/"

type SystemIke struct {
	DhGroup1          *SystemIkeDhGroup1  `json:"dh-group-1,omitempty"`
	DhGroup14         *SystemIkeDhGroup14 `json:"dh-group-14,omitempty"`
	DhGroup15         *SystemIkeDhGroup15 `json:"dh-group-15,omitempty"`
	DhGroup16         *SystemIkeDhGroup16 `json:"dh-group-16,omitempty"`
	DhGroup17         *SystemIkeDhGroup17 `json:"dh-group-17,omitempty"`
	DhGroup18         *SystemIkeDhGroup18 `json:"dh-group-18,omitempty"`
	DhGroup19         *SystemIkeDhGroup19 `json:"dh-group-19,omitempty"`
	DhGroup2          *SystemIkeDhGroup2  `json:"dh-group-2,omitempty"`
	DhGroup20         *SystemIkeDhGroup20 `json:"dh-group-20,omitempty"`
	DhGroup21         *SystemIkeDhGroup21 `json:"dh-group-21,omitempty"`
	DhGroup27         *SystemIkeDhGroup27 `json:"dh-group-27,omitempty"`
	DhGroup28         *SystemIkeDhGroup28 `json:"dh-group-28,omitempty"`
	DhGroup29         *SystemIkeDhGroup29 `json:"dh-group-29,omitempty"`
	DhGroup30         *SystemIkeDhGroup30 `json:"dh-group-30,omitempty"`
	DhGroup31         *SystemIkeDhGroup31 `json:"dh-group-31,omitempty"`
	DhGroup32         *SystemIkeDhGroup32 `json:"dh-group-32,omitempty"`
	DhGroup5          *SystemIkeDhGroup5  `json:"dh-group-5,omitempty"`
	DhKeypairCache    *string             `json:"dh-keypair-cache,omitempty"`
	DhKeypairCount    *int64              `json:"dh-keypair-count,omitempty"`
	DhKeypairThrottle *string             `json:"dh-keypair-throttle,omitempty"`
	DhMode            *string             `json:"dh-mode,omitempty"`
	DhMultiprocess    *string             `json:"dh-multiprocess,omitempty"`
	DhWorkerCount     *int64              `json:"dh-worker-count,omitempty"`
	EmbryonicLimit    *int64              `json:"embryonic-limit,omitempty"`
}

const SystemIkeDhGroup1Path = "system/ike/dh-group-1/"

type SystemIkeDhGroup1 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup14Path = "system/ike/dh-group-14/"

type SystemIkeDhGroup14 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup15Path = "system/ike/dh-group-15/"

type SystemIkeDhGroup15 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup16Path = "system/ike/dh-group-16/"

type SystemIkeDhGroup16 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup17Path = "system/ike/dh-group-17/"

type SystemIkeDhGroup17 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup18Path = "system/ike/dh-group-18/"

type SystemIkeDhGroup18 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup19Path = "system/ike/dh-group-19/"

type SystemIkeDhGroup19 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup2Path = "system/ike/dh-group-2/"

type SystemIkeDhGroup2 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup20Path = "system/ike/dh-group-20/"

type SystemIkeDhGroup20 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup21Path = "system/ike/dh-group-21/"

type SystemIkeDhGroup21 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup27Path = "system/ike/dh-group-27/"

type SystemIkeDhGroup27 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup28Path = "system/ike/dh-group-28/"

type SystemIkeDhGroup28 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup29Path = "system/ike/dh-group-29/"

type SystemIkeDhGroup29 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup30Path = "system/ike/dh-group-30/"

type SystemIkeDhGroup30 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup31Path = "system/ike/dh-group-31/"

type SystemIkeDhGroup31 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup32Path = "system/ike/dh-group-32/"

type SystemIkeDhGroup32 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}

const SystemIkeDhGroup5Path = "system/ike/dh-group-5/"

type SystemIkeDhGroup5 struct {
	KeypairCache *string `json:"keypair-cache,omitempty"`
	KeypairCount *int64  `json:"keypair-count,omitempty"`
	Mode         *string `json:"mode,omitempty"`
}
