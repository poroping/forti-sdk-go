package models

const EmailfilterFortishieldPath = "emailfilter/fortishield/"

type EmailfilterFortishield struct {
	SpamSubmitForce   *string `json:"spam-submit-force,omitempty"`
	SpamSubmitSrv     *string `json:"spam-submit-srv,omitempty"`
	SpamSubmitTxt2htm *string `json:"spam-submit-txt2htm,omitempty"`
}
