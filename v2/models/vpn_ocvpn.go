package models

const VpnOcvpnPath = "vpn/ocvpn/"

type VpnOcvpn struct {
	AutoDiscovery             *string                    `json:"auto-discovery,omitempty"`
	AutoDiscoveryShortcutMode *string                    `json:"auto-discovery-shortcut-mode,omitempty"`
	Eap                       *string                    `json:"eap,omitempty"`
	EapUsers                  *string                    `json:"eap-users,omitempty"`
	ForticlientAccess         *VpnOcvpnForticlientAccess `json:"forticlient-access,omitempty"`
	IpAllocationBlock         *string                    `json:"ip-allocation-block,omitempty"`
	Multipath                 *string                    `json:"multipath,omitempty"`
	Nat                       *string                    `json:"nat,omitempty"`
	Overlays                  *[]VpnOcvpnOverlays        `json:"overlays,omitempty"`
	PollInterval              *int64                     `json:"poll-interval,omitempty"`
	Role                      *string                    `json:"role,omitempty"`
	Sdwan                     *string                    `json:"sdwan,omitempty"`
	SdwanZone                 *string                    `json:"sdwan-zone,omitempty"`
	Status                    *string                    `json:"status,omitempty"`
	WanInterface              *[]VpnOcvpnWanInterface    `json:"wan-interface,omitempty"`
}

const VpnOcvpnForticlientAccessPath = "vpn/ocvpn/forticlient-access/"

type VpnOcvpnForticlientAccess struct {
	AuthGroups *[]VpnOcvpnForticlientAccessAuthGroups `json:"auth-groups,omitempty"`
	Psksecret  *string                                `json:"psksecret,omitempty"`
	Status     *string                                `json:"status,omitempty"`
}

const VpnOcvpnForticlientAccessAuthGroupsPath = "vpn/ocvpn/forticlient-access/auth-groups/"

type VpnOcvpnForticlientAccessAuthGroups struct {
	AuthGroup *string                                        `json:"auth-group,omitempty"`
	Name      *string                                        `json:"name,omitempty"`
	Overlays  *[]VpnOcvpnForticlientAccessAuthGroupsOverlays `json:"overlays,omitempty"`
}

const VpnOcvpnForticlientAccessAuthGroupsOverlaysPath = "vpn/ocvpn/forticlient-access/auth-groups/overlays/"

type VpnOcvpnForticlientAccessAuthGroupsOverlays struct {
	OverlayName *string `json:"overlay-name,omitempty"`
}

const VpnOcvpnOverlaysPath = "vpn/ocvpn/overlays/"

type VpnOcvpnOverlays struct {
	AssignIp     *string                    `json:"assign-ip,omitempty"`
	InterOverlay *string                    `json:"inter-overlay,omitempty"`
	Ipv4EndIp    *string                    `json:"ipv4-end-ip,omitempty"`
	Ipv4StartIp  *string                    `json:"ipv4-start-ip,omitempty"`
	OverlayName  *string                    `json:"overlay-name,omitempty"`
	Subnets      *[]VpnOcvpnOverlaysSubnets `json:"subnets,omitempty"`
}

const VpnOcvpnOverlaysSubnetsPath = "vpn/ocvpn/overlays/subnets/"

type VpnOcvpnOverlaysSubnets struct {
	Id        *int64  `json:"id,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Subnet    *string `json:"subnet,omitempty"`
	Type      *string `json:"type,omitempty"`
}

const VpnOcvpnWanInterfacePath = "vpn/ocvpn/wan-interface/"

type VpnOcvpnWanInterface struct {
	Name *string `json:"name,omitempty"`
}
