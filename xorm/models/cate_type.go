package models

import (
	"time"
)

type CateType struct {
	Id      int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Cateid  int       `json:"cateId" xorm:"not null comment('所属考试') INT(4)"`
	Title   string    `json:"title" xorm:"not null comment('类型标题') VARCHAR(100)"`
	Creater string    `json:"creater" xorm:"not null comment('创建人') VARCHAR(50)"`
	Time    time.Time `json:"time" xorm:"not null comment('创建时间') DATETIME"`
	Order   int       `json:"order" xorm:"not null comment('排序') TINYINT(3)"`
	Type    int       `json:"type" xorm:"not null default 0 comment('类型0课程1合集2自动上下线课程') TINYINT(4)"`
	Rid     int       `json:"rid" xorm:"not null comment('合集/课程Id') INT(10)"`
	Typeid  int       `json:"typeId" xorm:"not null comment('类型Id') INT(10)"`
	Img     string    `json:"img" xorm:"not null default '' comment('图片') VARCHAR(100)"`
	Number  int       `json:"number" xorm:"not null default 5 comment('展示条数') TINYINT(3)"`
}
