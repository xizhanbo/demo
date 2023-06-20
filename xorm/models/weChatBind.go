package models

import (
	"time"
)

type Wechatbind struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Openid   string    `json:"openId" xorm:"not null default '' comment('微信OpenId') VARCHAR(100)"`
	Userid   int       `json:"userId" xorm:"not null comment('用户ID') INT(11)"`
	Username string    `json:"userName" xorm:"not null comment('用户名') VARCHAR(100)"`
	Binddate time.Time `json:"bindDate" xorm:"not null comment('绑定时间') DATETIME"`
	Nickname string    `json:"nickName" xorm:"not null comment('昵称') VARCHAR(100)"`
	Status   int       `json:"status" xorm:"not null default 1 comment('绑定状态') TINYINT(4)"`
	Miniid   string    `json:"miniId" xorm:"not null default '' comment('小程序openId') VARCHAR(100)"`
	Unionid  string    `json:"unionId" xorm:"not null default '' comment('微信unionI') VARCHAR(100)"`
}
