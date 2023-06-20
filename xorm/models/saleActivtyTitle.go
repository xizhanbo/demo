package models

import (
	"time"
)

type Saleactivtytitle struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name       string    `json:"name" xorm:"not null comment('活动名称') VARCHAR(255)"`
	Title      string    `json:"title" xorm:"not null comment('活动标题') VARCHAR(15)"`
	Subtitle   string    `json:"subtitle" xorm:"comment('活动副标题') VARCHAR(255)"`
	Status     int       `json:"status" xorm:"not null default 0 comment('是否上线：1上线，2下线，0未上线') TINYINT(4)"`
	Optname    string    `json:"optname" xorm:"comment('操作人') VARCHAR(255)"`
	Url        string    `json:"url" xorm:"comment('换购地址') VARCHAR(100)"`
	Uid        int       `json:"uid" xorm:"not null comment('用户id') INT(11)"`
	Createdate time.Time `json:"createDate" xorm:"comment('创建时间') DATETIME"`
	Editdate   time.Time `json:"editDate" xorm:"comment('编辑时间') DATETIME"`
}
