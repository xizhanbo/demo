package models

import (
	"time"
)

type TeacherLessonRelation struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TeacherLessonId int       `json:"teacher_lesson_id" xorm:"comment('教师网课件id') index INT(11)"`
	NetLessonId     int       `json:"net_lesson_id" xorm:"comment('在线课件id') index INT(11)"`
	Type            int       `json:"type" xorm:"default 0 comment('0录播 1直播') TINYINT(1)"`
	CreatedAt       time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"comment('修改时间') DATETIME"`
}
