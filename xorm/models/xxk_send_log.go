package models

import (
	"time"
)

type XxkSendLog struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ApplyOrder   string    `json:"apply_order" xorm:"not null default '0' comment('申请单编号') CHAR(20)"`
	CardMianzhi  int       `json:"card_mianzhi" xorm:"not null default 0 comment('学习卡面值') INT(11)"`
	Dailishang1  string    `json:"dailishang_1" xorm:"not null default '' comment('一级代理商') CHAR(20)"`
	CardNoStart  string    `json:"card_no_start" xorm:"not null default '0' comment('派送起始卡号') CHAR(11)"`
	CardNoEnd    string    `json:"card_no_end" xorm:"not null default '0' comment('派送结束卡号') CHAR(11)"`
	CardSendTime time.Time `json:"card_send_time" xorm:"not null default CURRENT_TIMESTAMP comment('派送学习卡的时间') TIMESTAMP"`
}
