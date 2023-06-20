package models

import (
	"time"
)

type ClassExercisesNum struct {
	Id                int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(10)"`
	ClassId           int       `json:"class_id" xorm:"not null comment('课件id') index INT(11)"`
	ClassExercisesNum int       `json:"class_exercises_num" xorm:"default 0 comment('随堂练习数量') index INT(11)"`
	AfterExercisesNum int       `json:"after_exercises_num" xorm:"default 0 comment('课后练习数量') index INT(11)"`
	CourseType        int       `json:"course_type" xorm:"default 0 comment('1录播2直播') index TINYINT(4)"`
	SubjectType       int       `json:"subject_type" xorm:"default 1 comment('课后练习科目类型 默认1行测2申论') index TINYINT(1)"`
	BuildType         int       `json:"build_type" xorm:"default -1 comment('试题类型1套题0单题') index TINYINT(1)"`
	CreatedAt         time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"DATETIME"`
}
