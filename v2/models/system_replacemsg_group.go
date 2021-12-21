package models

const SystemReplacemsgGroupPath = "system/replacemsg-group/"

type SystemReplacemsgGroup struct {
	Admin                 *[]SystemReplacemsgGroupAdmin                 `json:"admin,omitempty"`
	Alertmail             *[]SystemReplacemsgGroupAlertmail             `json:"alertmail,omitempty"`
	Auth                  *[]SystemReplacemsgGroupAuth                  `json:"auth,omitempty"`
	Automation            *[]SystemReplacemsgGroupAutomation            `json:"automation,omitempty"`
	Comment               *string                                       `json:"comment,omitempty"`
	CustomMessage         *[]SystemReplacemsgGroupCustomMessage         `json:"custom-message,omitempty"`
	DeviceDetectionPortal *[]SystemReplacemsgGroupDeviceDetectionPortal `json:"device-detection-portal,omitempty"`
	FortiguardWf          *[]SystemReplacemsgGroupFortiguardWf          `json:"fortiguard-wf,omitempty"`
	Ftp                   *[]SystemReplacemsgGroupFtp                   `json:"ftp,omitempty"`
	GroupType             *string                                       `json:"group-type,omitempty"`
	Http                  *[]SystemReplacemsgGroupHttp                  `json:"http,omitempty"`
	Icap                  *[]SystemReplacemsgGroupIcap                  `json:"icap,omitempty"`
	Mail                  *[]SystemReplacemsgGroupMail                  `json:"mail,omitempty"`
	NacQuar               *[]SystemReplacemsgGroupNacQuar               `json:"nac-quar,omitempty"`
	Name                  *string                                       `json:"name,omitempty"`
	Nntp                  *[]SystemReplacemsgGroupNntp                  `json:"nntp,omitempty"`
	Spam                  *[]SystemReplacemsgGroupSpam                  `json:"spam,omitempty"`
	Sslvpn                *[]SystemReplacemsgGroupSslvpn                `json:"sslvpn,omitempty"`
	TrafficQuota          *[]SystemReplacemsgGroupTrafficQuota          `json:"traffic-quota,omitempty"`
	Utm                   *[]SystemReplacemsgGroupUtm                   `json:"utm,omitempty"`
	Webproxy              *[]SystemReplacemsgGroupWebproxy              `json:"webproxy,omitempty"`
}

type SystemReplacemsgGroupAdmin struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupAlertmail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupAuth struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupAutomation struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupCustomMessage struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupDeviceDetectionPortal struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupFortiguardWf struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupFtp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupHttp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupIcap struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupMail struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupNacQuar struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupNntp struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupSpam struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupSslvpn struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupTrafficQuota struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupUtm struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}

type SystemReplacemsgGroupWebproxy struct {
	Buffer  *string `json:"buffer,omitempty"`
	Format  *string `json:"format,omitempty"`
	Header  *string `json:"header,omitempty"`
	MsgType *string `json:"msg-type,omitempty"`
}
