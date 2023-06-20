package models

import (
	"time"
)

type XxkCardNew struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CardNo        string    `json:"card_no" xorm:"not null default '0' comment('充值卡卡号') index CHAR(11)"`
	PwdEncrypt    string    `json:"pwd_encrypt" xorm:"not null default '0' comment('加密后的学习卡密码') CHAR(32)"`
	CardMianzhi   int       `json:"card_mianzhi" xorm:"not null default 0 comment('学习卡面值，现在有10，50，100，200，500这几种面值') INT(11)"`
	CardStatus    int       `json:"card_status" xorm:"not null default 0 comment('学习卡的状态 0:库存状态，1：已分配，2：已激活，3：已退回，4：已锁定') TINYINT(4)"`
	Dailishang1   string    `json:"dailishang_1" xorm:"comment('一级代理商') VARCHAR(20)"`
	Dailishang2   string    `json:"dailishang_2" xorm:"comment('二级代理商') VARCHAR(20)"`
	UsernameJihuo string    `json:"username_jihuo" xorm:"comment('使用学习卡的用户名') VARCHAR(20)"`
	JihuoTime     time.Time `json:"jihuo_time" xorm:"not null default '0000-00-00 00:00:00' comment('学习卡激活时间') TIMESTAMP"`
}
