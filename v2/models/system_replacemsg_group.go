package models

const SystemReplacemsgGroupPath = "system/replacemsg-group/"

type SystemReplacemsgGroup struct {
	Admin         *[]SystemReplacemsgGroupAdmin         `json:"admin,omitempty"`
	Alertmail     *[]SystemReplacemsgGroupAlertmail     `json:"alertmail,omitempty"`
	Auth          *[]SystemReplacemsgGroupAuth          `json:"auth,omitempty"`
	Automation    *[]SystemReplacemsgGroupAutomation    `json:"automation,omitempty"`
	Comment       *string                               `json:"comment,omitempty"`
	CustomMessage *[]SystemReplacemsgGroupCustomMessage `json:"custom-message,omitempty"`
	FortiguardWf  *[]SystemReplacemsgGroupFortiguardWf  `json:"fortiguard-wf,omitempty"`
	Ftp           *[]SystemReplacemsgGroupFtp           `json:"ftp,omitempty"`
	GroupType     *string                               `json:"group-type,omitempty"`
	Http          *[]SystemReplacemsgGroupHttp          `json:"http,omitempty"`
	Icap          *[]SystemReplacemsgGroupIcap          `json:"icap,omitempty"`
	Mail          *[]SystemReplacemsgGroupMail          `json:"mail,omitempty"`
	NacQuar       *[]SystemReplacemsgGroupNacQuar       `json:"nac-quar,omitempty"`
	Name          *string                               `json:"name,omitempty"`
	Spam          *[]SystemReplacemsgGroupSpam          `json:"spam,omitempty"`
	Sslvpn        *[]SystemReplacemsgGroupSslvpn        `json:"sslvpn,omitempty"`
	TrafficQuota  *[]SystemReplacemsgGroupTrafficQuota  `json:"traffic-quota,omitempty"`
	Utm           *[]SystemReplacemsgGroupUtm           `json:"utm,omitempty"`
	Webproxy      *[]SystemReplacemsgGroupWebproxy      `json:"webproxy,omitempty"`
}

const SystemReplacemsgGroupAdminPath = "system/replacemsg-group/admin/"

type SystemReplacemsgGroupAdmin struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupAlertmailPath = "system/replacemsg-group/alertmail/"

type SystemReplacemsgGroupAlertmail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupAuthPath = "system/replacemsg-group/auth/"

type SystemReplacemsgGroupAuth struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupAutomationPath = "system/replacemsg-group/automation/"

type SystemReplacemsgGroupAutomation struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupCustomMessagePath = "system/replacemsg-group/custom-message/"

type SystemReplacemsgGroupCustomMessage struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupFortiguardWfPath = "system/replacemsg-group/fortiguard-wf/"

type SystemReplacemsgGroupFortiguardWf struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupFtpPath = "system/replacemsg-group/ftp/"

type SystemReplacemsgGroupFtp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupHttpPath = "system/replacemsg-group/http/"

type SystemReplacemsgGroupHttp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupIcapPath = "system/replacemsg-group/icap/"

type SystemReplacemsgGroupIcap struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupMailPath = "system/replacemsg-group/mail/"

type SystemReplacemsgGroupMail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupNacQuarPath = "system/replacemsg-group/nac-quar/"

type SystemReplacemsgGroupNacQuar struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupSpamPath = "system/replacemsg-group/spam/"

type SystemReplacemsgGroupSpam struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupSslvpnPath = "system/replacemsg-group/sslvpn/"

type SystemReplacemsgGroupSslvpn struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupTrafficQuotaPath = "system/replacemsg-group/traffic-quota/"

type SystemReplacemsgGroupTrafficQuota struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupUtmPath = "system/replacemsg-group/utm/"

type SystemReplacemsgGroupUtm struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

const SystemReplacemsgGroupWebproxyPath = "system/replacemsg-group/webproxy/"

type SystemReplacemsgGroupWebproxy struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
