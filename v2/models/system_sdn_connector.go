package models

const SystemSdnConnectorPath = "system/sdn-connector/"

type SystemSdnConnector struct {
	AccessKey         *string                             `json:"access-key,omitempty"`
	ApiKey            *string                             `json:"api-key,omitempty"`
	AzureRegion       *string                             `json:"azure-region,omitempty"`
	ClientId          *string                             `json:"client-id,omitempty"`
	ClientSecret      *string                             `json:"client-secret,omitempty"`
	CompartmentId     *string                             `json:"compartment-id,omitempty"`
	ComputeGeneration *float64                            `json:"compute-generation,omitempty"`
	Domain            *string                             `json:"domain,omitempty"`
	ExternalIp        *[]SystemSdnConnectorExternalIp     `json:"external-ip,omitempty"`
	ForwardingRule    *[]SystemSdnConnectorForwardingRule `json:"forwarding-rule,omitempty"`
	GcpProject        *string                             `json:"gcp-project,omitempty"`
	GcpProjectList    *[]SystemSdnConnectorGcpProjectList `json:"gcp-project-list,omitempty"`
	GroupName         *string                             `json:"group-name,omitempty"`
	HaStatus          *string                             `json:"ha-status,omitempty"`
	IbmRegion         *string                             `json:"ibm-region,omitempty"`
	LoginEndpoint     *string                             `json:"login-endpoint,omitempty"`
	Name              *string                             `json:"name,omitempty"`
	Nic               *[]SystemSdnConnectorNic            `json:"nic,omitempty"`
	OciCert           *string                             `json:"oci-cert,omitempty"`
	OciFingerprint    *string                             `json:"oci-fingerprint,omitempty"`
	OciRegion         *string                             `json:"oci-region,omitempty"`
	OciRegionType     *string                             `json:"oci-region-type,omitempty"`
	Password          *string                             `json:"password,omitempty"`
	PrivateKey        *string                             `json:"private-key,omitempty"`
	Region            *string                             `json:"region,omitempty"`
	ResourceGroup     *string                             `json:"resource-group,omitempty"`
	ResourceUrl       *string                             `json:"resource-url,omitempty"`
	Route             *[]SystemSdnConnectorRoute          `json:"route,omitempty"`
	RouteTable        *[]SystemSdnConnectorRouteTable     `json:"route-table,omitempty"`
	SecretKey         *string                             `json:"secret-key,omitempty"`
	SecretToken       *string                             `json:"secret-token,omitempty"`
	Server            *string                             `json:"server,omitempty"`
	ServerList        *[]SystemSdnConnectorServerList     `json:"server-list,omitempty"`
	ServerPort        *float64                            `json:"server-port,omitempty"`
	ServiceAccount    *string                             `json:"service-account,omitempty"`
	Status            *string                             `json:"status,omitempty"`
	SubscriptionId    *string                             `json:"subscription-id,omitempty"`
	TenantId          *string                             `json:"tenant-id,omitempty"`
	Type              *string                             `json:"type,omitempty"`
	UpdateInterval    *float64                            `json:"update-interval,omitempty"`
	UseMetadataIam    *string                             `json:"use-metadata-iam,omitempty"`
	UserId            *string                             `json:"user-id,omitempty"`
	Username          *string                             `json:"username,omitempty"`
	VcenterPassword   *string                             `json:"vcenter-password,omitempty"`
	VcenterServer     *string                             `json:"vcenter-server,omitempty"`
	VcenterUsername   *string                             `json:"vcenter-username,omitempty"`
	VerifyCertificate *string                             `json:"verify-certificate,omitempty"`
	VpcId             *string                             `json:"vpc-id,omitempty"`
}

type SystemSdnConnectorExternalIp struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdnConnectorForwardingRule struct {
	RuleName *string `json:"rule-name,omitempty"`
	Target   *string `json:"target,omitempty"`
}

type SystemSdnConnectorGcpProjectList struct {
	GcpZoneList *[]SystemSdnConnectorGcpProjectListGcpZoneList `json:"gcp-zone-list,omitempty"`
	Id          *string                                        `json:"id,omitempty"`
}

type SystemSdnConnectorGcpProjectListGcpZoneList struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdnConnectorNic struct {
	Ip   *[]SystemSdnConnectorNicIp `json:"ip,omitempty"`
	Name *string                    `json:"name,omitempty"`
}

type SystemSdnConnectorNicIp struct {
	Name          *string `json:"name,omitempty"`
	PublicIp      *string `json:"public-ip,omitempty"`
	ResourceGroup *string `json:"resource-group,omitempty"`
}

type SystemSdnConnectorRoute struct {
	Name *string `json:"name,omitempty"`
}

type SystemSdnConnectorRouteTable struct {
	Name           *string                              `json:"name,omitempty"`
	ResourceGroup  *string                              `json:"resource-group,omitempty"`
	Route          *[]SystemSdnConnectorRouteTableRoute `json:"route,omitempty"`
	SubscriptionId *string                              `json:"subscription-id,omitempty"`
}

type SystemSdnConnectorRouteTableRoute struct {
	Name    *string `json:"name,omitempty"`
	NextHop *string `json:"next-hop,omitempty"`
}

type SystemSdnConnectorServerList struct {
	Ip *string `json:"ip,omitempty"`
}
