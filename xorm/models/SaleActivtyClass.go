package models

import (
	"time"
)

type Saleactivtyclass struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Tid        int       `json:"tid" xorm:"not null comment('标题id') INT(11)"`
	Sid        int       `json:"sid" xorm:"not null comment('活动id') INT(11)"`
	Classid    int       `json:"classid" xorm:"not null comment('课程id') unique INT(11)"`
	Status     int       `json:"status" xorm:"not null default 0 comment('状态：0未上线，1上线，2下线，4删除') TINYINT(4)"`
	Createdate time.Time `json:"createDate" xorm:"comment('创建时间') DATETIME"`
	Editdate   time.Time `json:"editDate" xorm:"comment('编辑时间') DATETIME"`
}
