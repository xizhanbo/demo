package models

import (
	"time"
)

type Autoonlinecollection struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Rid        int       `json:"rid" xorm:"not null comment('添加进集合的课程ID，主键') index INT(10)"`
	Pid        int       `json:"Pid" xorm:"not null comment('集合ID') index INT(10)"`
	Status     int       `json:"Status" xorm:"not null TINYINT(4)"`
	Createdate time.Time `json:"CreateDate" xorm:"not null comment('创建时间') DATETIME"`
	Optuser    string    `json:"OptUser" xorm:"not null comment('操作人') VARCHAR(50)"`
	Title      string    `json:"Title" xorm:"not null comment('集合标题') VARCHAR(255)"`
	Statusdate time.Time `json:"StatusDate" xorm:"comment('更新时间（记录第一次上线时间）') DATETIME"`
}
