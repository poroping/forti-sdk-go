package models

const VpnOcvpnPath = "vpn/ocvpn/"

type VpnOcvpn struct {
	AutoDiscovery             *string                      `json:"auto-discovery,omitempty"`
	AutoDiscoveryShortcutMode *string                      `json:"auto-discovery-shortcut-mode,omitempty"`
	Eap                       *string                      `json:"eap,omitempty"`
	EapUsers                  *string                      `json:"eap-users,omitempty"`
	ForticlientAccess         *[]VpnOcvpnForticlientAccess `json:"forticlient-access,omitempty"`
	IpAllocationBlock         *string                      `json:"ip-allocation-block,omitempty"`
	Multipath                 *string                      `json:"multipath,omitempty"`
	Nat                       *string                      `json:"nat,omitempty"`
	Overlays                  *[]VpnOcvpnOverlays          `json:"overlays,omitempty"`
	PollInterval              *float64                     `json:"poll-interval,omitempty"`
	Role                      *string                      `json:"role,omitempty"`
	Sdwan                     *string                      `json:"sdwan,omitempty"`
	SdwanZone                 *string                      `json:"sdwan-zone,omitempty"`
	Status                    *string                      `json:"status,omitempty"`
	WanInterface              *[]VpnOcvpnWanInterface      `json:"wan-interface,omitempty"`
}

type VpnOcvpnForticlientAccess struct {
	AuthGroups *[]VpnOcvpnForticlientAccessAuthGroups `json:"auth-groups,omitempty"`
	Psksecret  *string                                `json:"psksecret,omitempty"`
	Status     *string                                `json:"status,omitempty"`
}

type VpnOcvpnForticlientAccessAuthGroups struct {
	AuthGroup *string                                        `json:"auth-group,omitempty"`
	Name      *string                                        `json:"name,omitempty"`
	Overlays  *[]VpnOcvpnForticlientAccessAuthGroupsOverlays `json:"overlays,omitempty"`
}

type VpnOcvpnForticlientAccessAuthGroupsOverlays struct {
	OverlayName *string `json:"overlay-name,omitempty"`
}

type VpnOcvpnOverlays struct {
	AssignIp     *string                    `json:"assign-ip,omitempty"`
	Id           *float64                   `json:"id,omitempty"`
	InterOverlay *string                    `json:"inter-overlay,omitempty"`
	Ipv4EndIp    *string                    `json:"ipv4-end-ip,omitempty"`
	Ipv4StartIp  *string                    `json:"ipv4-start-ip,omitempty"`
	Name         *string                    `json:"name,omitempty"`
	OverlayName  *string                    `json:"overlay-name,omitempty"`
	Subnets      *[]VpnOcvpnOverlaysSubnets `json:"subnets,omitempty"`
}

type VpnOcvpnOverlaysSubnets struct {
	Id        *float64 `json:"id,omitempty"`
	Interface *string  `json:"interface,omitempty"`
	Subnet    *string  `json:"subnet,omitempty"`
	Type      *string  `json:"type,omitempty"`
}

type VpnOcvpnWanInterface struct {
	Name *string `json:"name,omitempty"`
}
