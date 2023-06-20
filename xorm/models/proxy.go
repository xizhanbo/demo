package models

import (
	"time"
)

type Proxy struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	RealName         string    `json:"real_name" xorm:"default '' comment('代理身份证姓名') VARCHAR(50)"`
	CardNo           string    `json:"card_no" xorm:"default '' comment('身份证号') unique CHAR(18)"`
	UserMobile       string    `json:"user_mobile" xorm:"not null comment('用户手机号') unique CHAR(11)"`
	BankNo           string    `json:"bank_no" xorm:"default '' comment('银行卡号') VARCHAR(20)"`
	CardNoFront      string    `json:"card_no_front" xorm:"default '' comment('身份证正面照') VARCHAR(100)"`
	CardNoBack       string    `json:"card_no_back" xorm:"default '' comment('身份证反面照') VARCHAR(100)"`
	CardNoHand       string    `json:"card_no_hand" xorm:"default '' comment('手持身份证照') VARCHAR(100)"`
	Bank             string    `json:"bank" xorm:"default '' comment('开户行') VARCHAR(100)"`
	BankBranch       string    `json:"bank_branch" xorm:"default '' comment('所在支行') VARCHAR(100)"`
	BankArea         string    `json:"bank_area" xorm:"default '' comment('开户地区') VARCHAR(100)"`
	UserStatus       int       `json:"user_status" xorm:"not null default 0 comment('0 已注冊 1待审核 2审核未通过 3审核通过 4关闭 5锁定') index TINYINT(1)"`
	ParentId         int       `json:"parent_id" xorm:"not null default 0 comment('相对应父级') index INT(11)"`
	AllyMoney        string    `json:"ally_money" xorm:"not null default 0.00 comment('盟友贡献总金额') DECIMAL(10,2)"`
	NickName         string    `json:"nick_name" xorm:"not null comment('微信昵称') VARCHAR(50)"`
	CreateTime       time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	CurrentResidence string    `json:"current_residence" xorm:"not null default '' comment('现住地') VARCHAR(100)"`
	Directsales      string    `json:"directSales" xorm:"not null default 0.00 comment('个人直售总金额') DECIMAL(10,2)"`
	Password         string    `json:"password" xorm:"not null default '' comment('合伙人密码') VARCHAR(100)"`
	Unionid          string    `json:"unionId" xorm:"not null comment('微信unionId') VARCHAR(100)"`
	LockTime         time.Time `json:"lock_time" xorm:"comment('锁定时间') DATETIME"`
	PersonalCode     string    `json:"personal_code" xorm:"default '' comment('个人直售二维码') VARCHAR(100)"`
	ProxyCode        string    `json:"proxy_code" xorm:"default '' comment('盟友贡献二维码') VARCHAR(100)"`
}
