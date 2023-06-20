package models

import (
	"time"
)

type AutoCloseCollection struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Title     string    `json:"title" xorm:"comment('集合名称') VARCHAR(100)"`
	Limit     int       `json:"limit" xorm:"not null default 0 comment('限报人数') INT(10)"`
	Status    int       `json:"status" xorm:"not null default 0 comment('0 未上线  1 已上线 2已下线') TINYINT(1)"`
	Creater   string    `json:"creater" xorm:"not null comment('创建人') VARCHAR(50)"`
	Pid       int       `json:"pid" xorm:"not null default 0 comment('父级id') index INT(10)"`
	Rid       int       `json:"rid" xorm:"not null default 0 comment('课程id') index INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
