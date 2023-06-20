package models

import (
	"time"
)

type TeacherTreaty struct {
	TeacherAgreeId int       `json:"teacher_agree_id" xorm:"not null comment('教师网协议Id') index INT(11)"`
	NetTreatyId    int       `json:"net_treaty_id" xorm:"not null comment('在线协议Id') index INT(11)"`
	CreatedAt      time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"comment('修改时间') DATETIME"`
}
