package models

import (
	"time"
)

type UnfinishExamStatus struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SyllabusId int       `json:"syllabus_id" xorm:"not null comment('大纲节点id') index index(syllabus_user) INT(11)"`
	NetClassId int       `json:"net_class_id" xorm:"not null comment('课程id') index INT(11)"`
	Status     int       `json:"status" xorm:"comment('0未读1已读') TINYINT(1)"`
	UserName   string    `json:"user_name" xorm:"not null comment('用户名') index(syllabus_user) VARCHAR(150)"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DATETIME"`
}
