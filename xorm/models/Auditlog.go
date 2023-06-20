package models

import (
	"time"
)

type Auditlog struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	AuditId          int       `json:"audit_id" xorm:"not null comment('审核表中自增Id') index INT(1)"`
	ProxyId          int       `json:"proxy_id" xorm:"not null INT(11)"`
	Status           int       `json:"status" xorm:"not null comment('审核状态') index TINYINT(1)"`
	RealName         string    `json:"real_name" xorm:"comment('真实姓名') VARCHAR(20)"`
	Phone            string    `json:"phone" xorm:"not null comment('手机号	') CHAR(11)"`
	BankArea         string    `json:"bank_area" xorm:"comment('开户地区') VARCHAR(100)"`
	CardNo           string    `json:"card_no" xorm:"comment('身份证号') VARCHAR(18)"`
	CardNoFront      string    `json:"card_no_front" xorm:"comment('身份证正面图') VARCHAR(200)"`
	CardNoBack       string    `json:"card_no_back" xorm:"comment('身份证反面图') VARCHAR(200)"`
	CardNoHand       string    `json:"card_no_hand" xorm:"comment('手持身份证图') VARCHAR(200)"`
	BankNo           string    `json:"bank_no" xorm:"comment('银行卡号') VARCHAR(20)"`
	Bank             string    `json:"bank" xorm:"comment('开户行') VARCHAR(100)"`
	BankBranch       string    `json:"bank_branch" xorm:"not null comment('所在支行') VARCHAR(100)"`
	CreateTime       time.Time `json:"create_time" xorm:"not null comment('审核时间') DATETIME"`
	Memo             string    `json:"memo" xorm:"comment('审核备注') TEXT"`
	AuditorName      string    `json:"auditor_name" xorm:"not null comment('审核操作人') VARCHAR(20)"`
	NickName         string    `json:"nick_name" xorm:"not null comment('微信昵称') VARCHAR(20)"`
	CurrentResidence string    `json:"current_residence" xorm:"not null comment('现住地') VARCHAR(100)"`
}
