package models

import (
	"time"
)

type SyllabusTrialSwitch struct {
	SyllabusId int       `json:"syllabus_id" xorm:"not null pk comment('大纲节点id') INT(11)"`
	NetClassId int       `json:"net_class_id" xorm:"comment('课程id') index INT(11)"`
	OnOff      int       `json:"on_off" xorm:"comment('是否展示试听课件  0不展示') TINYINT(1)"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DATETIME"`
}
