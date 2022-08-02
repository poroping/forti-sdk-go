package models

const UserGroupPath = "user/group/"

type UserGroup struct {
	AuthConcurrentOverride *string            `json:"auth-concurrent-override,omitempty"`
	AuthConcurrentValue    *int64             `json:"auth-concurrent-value,omitempty"`
	Authtimeout            *int64             `json:"authtimeout,omitempty"`
	Company                *string            `json:"company,omitempty"`
	Email                  *string            `json:"email,omitempty"`
	Expire                 *int64             `json:"expire,omitempty"`
	ExpireType             *string            `json:"expire-type,omitempty"`
	GroupType              *string            `json:"group-type,omitempty"`
	Guest                  *[]UserGroupGuest  `json:"guest,omitempty"`
	HttpDigestRealm        *string            `json:"http-digest-realm,omitempty"`
	Id                     *int64             `json:"id,omitempty"`
	Match                  *[]UserGroupMatch  `json:"match,omitempty"`
	MaxAccounts            *int64             `json:"max-accounts,omitempty"`
	Member                 *[]UserGroupMember `json:"member,omitempty"`
	MobilePhone            *string            `json:"mobile-phone,omitempty"`
	MultipleGuestAdd       *string            `json:"multiple-guest-add,omitempty"`
	Name                   *string            `json:"name,omitempty"`
	Password               *string            `json:"password,omitempty"`
	SmsCustomServer        *string            `json:"sms-custom-server,omitempty"`
	SmsServer              *string            `json:"sms-server,omitempty"`
	Sponsor                *string            `json:"sponsor,omitempty"`
	SsoAttributeValue      *string            `json:"sso-attribute-value,omitempty"`
	UserId                 *string            `json:"user-id,omitempty"`
	UserName               *string            `json:"user-name,omitempty"`
}

const UserGroupGuestPath = "user/group/guest/"

type UserGroupGuest struct {
	Comment     *string `json:"comment,omitempty"`
	Company     *string `json:"company,omitempty"`
	Email       *string `json:"email,omitempty"`
	Expiration  *string `json:"expiration,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	MobilePhone *string `json:"mobile-phone,omitempty"`
	Name        *string `json:"name,omitempty"`
	Password    *string `json:"password,omitempty"`
	Sponsor     *string `json:"sponsor,omitempty"`
	UserId      *string `json:"user-id,omitempty"`
}

const UserGroupMatchPath = "user/group/match/"

type UserGroupMatch struct {
	GroupName  *string `json:"group-name,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	ServerName *string `json:"server-name,omitempty"`
}

const UserGroupMemberPath = "user/group/member/"

type UserGroupMember struct {
	Name *string `json:"name,omitempty"`
}
