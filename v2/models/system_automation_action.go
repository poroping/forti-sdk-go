package models

const SystemAutomationActionPath = "system/automation-action/"

type SystemAutomationAction struct {
	Accprofile                    *string                               `json:"accprofile,omitempty"`
	ActionType                    *string                               `json:"action-type,omitempty"`
	AlicloudAccessKeyId           *string                               `json:"alicloud-access-key-id,omitempty"`
	AlicloudAccessKeySecret       *string                               `json:"alicloud-access-key-secret,omitempty"`
	AlicloudAccountId             *string                               `json:"alicloud-account-id,omitempty"`
	AlicloudFunction              *string                               `json:"alicloud-function,omitempty"`
	AlicloudFunctionAuthorization *string                               `json:"alicloud-function-authorization,omitempty"`
	AlicloudFunctionDomain        *string                               `json:"alicloud-function-domain,omitempty"`
	AlicloudRegion                *string                               `json:"alicloud-region,omitempty"`
	AlicloudService               *string                               `json:"alicloud-service,omitempty"`
	AlicloudVersion               *string                               `json:"alicloud-version,omitempty"`
	AwsApiId                      *string                               `json:"aws-api-id,omitempty"`
	AwsApiKey                     *string                               `json:"aws-api-key,omitempty"`
	AwsApiPath                    *string                               `json:"aws-api-path,omitempty"`
	AwsApiStage                   *string                               `json:"aws-api-stage,omitempty"`
	AwsDomain                     *string                               `json:"aws-domain,omitempty"`
	AwsRegion                     *string                               `json:"aws-region,omitempty"`
	AzureApiKey                   *string                               `json:"azure-api-key,omitempty"`
	AzureApp                      *string                               `json:"azure-app,omitempty"`
	AzureDomain                   *string                               `json:"azure-domain,omitempty"`
	AzureFunction                 *string                               `json:"azure-function,omitempty"`
	AzureFunctionAuthorization    *string                               `json:"azure-function-authorization,omitempty"`
	Delay                         *int64                                `json:"delay,omitempty"`
	Description                   *string                               `json:"description,omitempty"`
	EmailFrom                     *string                               `json:"email-from,omitempty"`
	EmailSubject                  *string                               `json:"email-subject,omitempty"`
	EmailTo                       *[]SystemAutomationActionEmailTo      `json:"email-to,omitempty"`
	ExecuteSecurityFabric         *string                               `json:"execute-security-fabric,omitempty"`
	GcpFunction                   *string                               `json:"gcp-function,omitempty"`
	GcpFunctionDomain             *string                               `json:"gcp-function-domain,omitempty"`
	GcpFunctionRegion             *string                               `json:"gcp-function-region,omitempty"`
	GcpProject                    *string                               `json:"gcp-project,omitempty"`
	Headers                       *[]SystemAutomationActionHeaders      `json:"headers,omitempty"`
	HttpBody                      *string                               `json:"http-body,omitempty"`
	Message                       *string                               `json:"message,omitempty"`
	MessageType                   *string                               `json:"message-type,omitempty"`
	Method                        *string                               `json:"method,omitempty"`
	MinimumInterval               *int64                                `json:"minimum-interval,omitempty"`
	Name                          *string                               `json:"name,omitempty"`
	Port                          *int64                                `json:"port,omitempty"`
	Protocol                      *string                               `json:"protocol,omitempty"`
	ReplacementMessage            *string                               `json:"replacement-message,omitempty"`
	ReplacemsgGroup               *string                               `json:"replacemsg-group,omitempty"`
	Required                      *string                               `json:"required,omitempty"`
	Script                        *string                               `json:"script,omitempty"`
	SdnConnector                  *[]SystemAutomationActionSdnConnector `json:"sdn-connector,omitempty"`
	SecurityTag                   *string                               `json:"security-tag,omitempty"`
	TlsCertificate                *string                               `json:"tls-certificate,omitempty"`
	Uri                           *string                               `json:"uri,omitempty"`
	VerifyHostCert                *string                               `json:"verify-host-cert,omitempty"`
}

const SystemAutomationActionEmailToPath = "system/automation-action/email-to/"

type SystemAutomationActionEmailTo struct {
	Name *string `json:"name,omitempty"`
}

const SystemAutomationActionHeadersPath = "system/automation-action/headers/"

type SystemAutomationActionHeaders struct {
	Header *string `json:"header,omitempty"`
}

const SystemAutomationActionSdnConnectorPath = "system/automation-action/sdn-connector/"

type SystemAutomationActionSdnConnector struct {
	Name *string `json:"name,omitempty"`
}
