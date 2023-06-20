package models

import (
	"time"
)

type StudyReportSwitch struct {
	SyllabusId int       `json:"syllabus_id" xorm:"not null pk comment('大纲节点id') INT(11)"`
	OnOff      int       `json:"on_off" xorm:"default 0 comment('学习报告0关闭1开启') TINYINT(1)"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DATETIME"`
}
