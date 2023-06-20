package models

import (
	"time"
)

type LiveClass struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Roomid    string    `json:"roomId" xorm:"not null VARCHAR(20)"`
	Rid       int       `json:"rid" xorm:"not null INT(11)"`
	Type      int       `json:"type" xorm:"not null default 0 comment('类型0课程1合集2自动上下线') TINYINT(4)"`
	Creater   string    `json:"creater" xorm:"not null VARCHAR(50)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	IsDel     int       `json:"is_del" xorm:"not null default 0 comment('0 否 1是') TINYINT(1)"`
	Status    int       `json:"status" xorm:"not null default 0 comment('0 未上架 1 上架') TINYINT(1)"`
	SendTime  time.Time `json:"send_time" xorm:"comment('推送时间') DATETIME"`
	Order     int       `json:"order" xorm:"not null TINYINT(3)"`
}
