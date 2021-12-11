package models

const SystemAutomationActionPath = "system/automation-action/"

type SystemAutomationAction struct {
	Accprofile                    *string                              `json:"accprofile,omitempty"`
	ActionType                    *string                              `json:"action-type,omitempty"`
	AlicloudAccessKeyId           *string                              `json:"alicloud-access-key-id,omitempty"`
	AlicloudAccessKeySecret       *string                              `json:"alicloud-access-key-secret,omitempty"`
	AlicloudFunctionAuthorization *string                              `json:"alicloud-function-authorization,omitempty"`
	AwsApiKey                     *string                              `json:"aws-api-key,omitempty"`
	AzureApiKey                   *string                              `json:"azure-api-key,omitempty"`
	AzureFunctionAuthorization    *string                              `json:"azure-function-authorization,omitempty"`
	Description                   *string                              `json:"description,omitempty"`
	EmailFrom                     *string                              `json:"email-from,omitempty"`
	EmailSubject                  *string                              `json:"email-subject,omitempty"`
	EmailTo                       []SystemAutomationActionEmailTo      `json:"email-to,omitempty"`
	ExecuteSecurityFabric         *string                              `json:"execute-security-fabric,omitempty"`
	Headers                       []SystemAutomationActionHeaders      `json:"headers,omitempty"`
	HttpBody                      *string                              `json:"http-body,omitempty"`
	Message                       *string                              `json:"message,omitempty"`
	MessageType                   *string                              `json:"message-type,omitempty"`
	Method                        *string                              `json:"method,omitempty"`
	MinimumInterval               *int64                               `json:"minimum-interval,omitempty"`
	Name                          *string                              `json:"name,omitempty"`
	Port                          *int64                               `json:"port,omitempty"`
	Protocol                      *string                              `json:"protocol,omitempty"`
	ReplacementMessage            *string                              `json:"replacement-message,omitempty"`
	ReplacemsgGroup               *string                              `json:"replacemsg-group,omitempty"`
	Script                        *string                              `json:"script,omitempty"`
	SdnConnector                  []SystemAutomationActionSdnConnector `json:"sdn-connector,omitempty"`
	SecurityTag                   *string                              `json:"security-tag,omitempty"`
	TlsCertificate                *string                              `json:"tls-certificate,omitempty"`
	Uri                           *string                              `json:"uri,omitempty"`
	VerifyHostCert                *string                              `json:"verify-host-cert,omitempty"`
}

type SystemAutomationActionEmailTo struct {
	Name *string `json:"name,omitempty"`
}

type SystemAutomationActionHeaders struct {
	Header *string `json:"header,omitempty"`
}

type SystemAutomationActionSdnConnector struct {
	Name *string `json:"name,omitempty"`
}
