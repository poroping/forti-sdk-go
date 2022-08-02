package models

const SystemWccpPath = "system/wccp/"

type SystemWccp struct {
	AssignmentBucketFormat *string `json:"assignment-bucket-format,omitempty"`
	AssignmentDstaddrMask  *string `json:"assignment-dstaddr-mask,omitempty"`
	AssignmentMethod       *string `json:"assignment-method,omitempty"`
	AssignmentSrcaddrMask  *string `json:"assignment-srcaddr-mask,omitempty"`
	AssignmentWeight       *int64  `json:"assignment-weight,omitempty"`
	Authentication         *string `json:"authentication,omitempty"`
	CacheEngineMethod      *string `json:"cache-engine-method,omitempty"`
	CacheId                *string `json:"cache-id,omitempty"`
	ForwardMethod          *string `json:"forward-method,omitempty"`
	GroupAddress           *string `json:"group-address,omitempty"`
	Password               *string `json:"password,omitempty"`
	Ports                  *string `json:"ports,omitempty"`
	PortsDefined           *string `json:"ports-defined,omitempty"`
	PrimaryHash            *string `json:"primary-hash,omitempty"`
	Priority               *int64  `json:"priority,omitempty"`
	Protocol               *int64  `json:"protocol,omitempty"`
	ReturnMethod           *string `json:"return-method,omitempty"`
	RouterId               *string `json:"router-id,omitempty"`
	RouterList             *string `json:"router-list,omitempty"`
	ServerList             *string `json:"server-list,omitempty"`
	ServerType             *string `json:"server-type,omitempty"`
	ServiceId              *string `json:"service-id,omitempty"`
	ServiceType            *string `json:"service-type,omitempty"`
}
