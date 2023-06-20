package models

import (
	"time"
)

type Partnerpersonal struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username   string    `json:"userName" xorm:"not null default '' comment('学员用户名') VARCHAR(80)"`
	Proxyid    int       `json:"proxyId" xorm:"not null comment('代理id') index INT(11)"`
	Createtime time.Time `json:"createTime" xorm:"not null comment('添加时间') DATETIME"`
}
