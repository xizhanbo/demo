package models

import (
	"time"
)

type StageExamQuestions struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ExamQuestionId int       `json:"exam_question_id" xorm:"not null comment('阶段测试题id') index INT(11)"`
	Name           string    `json:"name" xorm:"not null comment('阶段测试题名称') index CHAR(200)"`
	IsLongTerm     int       `json:"is_long_term" xorm:"default 0 comment('考试时间是否有效0否1是') TINYINT(1)"`
	Status         int       `json:"status" xorm:"comment('1正常') TINYINT(1)"`
	StartTime      time.Time `json:"start_time" xorm:"comment('开始时间') index DATETIME"`
	EndTime        time.Time `json:"end_time" xorm:"comment('结束时间') index DATETIME"`
	CreatedAt      time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"DATETIME"`
}
