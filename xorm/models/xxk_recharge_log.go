package models

import (
	"time"
)

type XxkRechargeLog struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CardNo         string    `json:"card_no" xorm:"not null default '0' comment('学习卡卡号') CHAR(20)"`
	RechargeStatus int       `json:"recharge_status" xorm:"not null default 0 comment('学习卡状态，0：锁定状态（正在充值），1：充值成功') TINYINT(4)"`
	RechargeUser   string    `json:"recharge_user" xorm:"not null default '' comment('充值的用户') index VARCHAR(20)"`
	RechargeTime   time.Time `json:"recharge_time" xorm:"not null default CURRENT_TIMESTAMP comment('充值时间') TIMESTAMP"`
}
