package models

const SystemCsfPath = "system/csf/"

type SystemCsf struct {
	AcceptAuthByCert         *string                     `json:"accept-auth-by-cert,omitempty"`
	AuthorizationRequestType *string                     `json:"authorization-request-type,omitempty"`
	Certificate              *string                     `json:"certificate,omitempty"`
	ConfigurationSync        *string                     `json:"configuration-sync,omitempty"`
	DownstreamAccess         *string                     `json:"downstream-access,omitempty"`
	DownstreamAccprofile     *string                     `json:"downstream-accprofile,omitempty"`
	FabricConnector          *[]SystemCsfFabricConnector `json:"fabric-connector,omitempty"`
	FabricDevice             *[]SystemCsfFabricDevice    `json:"fabric-device,omitempty"`
	FabricObjectUnification  *string                     `json:"fabric-object-unification,omitempty"`
	FabricWorkers            *float64                    `json:"fabric-workers,omitempty"`
	GroupName                *string                     `json:"group-name,omitempty"`
	GroupPassword            *string                     `json:"group-password,omitempty"`
	LogUnification           *string                     `json:"log-unification,omitempty"`
	ManagementIp             *string                     `json:"management-ip,omitempty"`
	ManagementPort           *float64                    `json:"management-port,omitempty"`
	SamlConfigurationSync    *string                     `json:"saml-configuration-sync,omitempty"`
	Status                   *string                     `json:"status,omitempty"`
	TrustedList              *[]SystemCsfTrustedList     `json:"trusted-list,omitempty"`
	Upstream                 *string                     `json:"upstream,omitempty"`
	UpstreamIp               *string                     `json:"upstream-ip,omitempty"`
	UpstreamPort             *float64                    `json:"upstream-port,omitempty"`
}

type SystemCsfFabricConnector struct {
	Accprofile               *string `json:"accprofile,omitempty"`
	ConfigurationWriteAccess *string `json:"configuration-write-access,omitempty"`
	Serial                   *string `json:"serial,omitempty"`
}

type SystemCsfFabricDevice struct {
	AccessToken *string  `json:"access-token,omitempty"`
	DeviceIp    *string  `json:"device-ip,omitempty"`
	HttpsPort   *float64 `json:"https-port,omitempty"`
	Name        *string  `json:"name,omitempty"`
}

type SystemCsfTrustedList struct {
	Action                  *string `json:"action,omitempty"`
	AuthorizationType       *string `json:"authorization-type,omitempty"`
	Certificate             *string `json:"certificate,omitempty"`
	DownstreamAuthorization *string `json:"downstream-authorization,omitempty"`
	HaMembers               *string `json:"ha-members,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Serial                  *string `json:"serial,omitempty"`
}
