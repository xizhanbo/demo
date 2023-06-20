package models

import (
	"time"
)

type TeacherClassRelation struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TeacherClassId int       `json:"teacher_class_id" xorm:"not null comment('教师网课程id') index INT(11)"`
	NetClassId     int       `json:"net_class_id" xorm:"not null comment('在线课程id') index INT(11)"`
	CreatedAt      time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"comment('修改时间') DATETIME"`
}
