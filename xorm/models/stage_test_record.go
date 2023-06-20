package models

import (
	"time"
)

type StageTestRecord struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	NetClassId int       `json:"net_class_id" xorm:"not null comment('课程id') index INT(11)"`
	SyllabusId int       `json:"syllabus_id" xorm:"not null comment('大纲id') index INT(11)"`
	UserName   string    `json:"user_name" xorm:"not null comment('用户名') index CHAR(200)"`
	IsFinish   int       `json:"is_finish" xorm:"comment('是否完成0否1是') TINYINT(1)"`
	Origin     int       `json:"origin" xorm:"not null comment('1-华图在线, 2-面库') TINYINT(1)"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DATETIME"`
}
