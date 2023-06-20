package models

import (
	"time"
)

type ParnterCate struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Rid       int       `json:"rid" xorm:"not null comment('课程/合集id') INT(10)"`
	Type      int       `json:"type" xorm:"not null default 0 comment('类型0课程1合集2自动上下线') TINYINT(4)"`
	Order     int       `json:"order" xorm:"not null comment('排序') TINYINT(3)"`
	IsDel     int       `json:"is_del" xorm:"not null default 0 comment('是否删除 0 否 1是') TINYINT(2)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	Creater   string    `json:"creater" xorm:"not null comment('创建人') VARCHAR(100)"`
	Pid       int       `json:"pid" xorm:"not null comment('合作方id') INT(10)"`
}
