package models

import (
	"time"
)

type ZhuokunTeacherForOnlineTearchers struct {
	Id                 int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ZhuokunTeacherId   int64     `json:"zhuokun_teacher_id" xorm:"comment('卓坤医疗老师ID') index BIGINT(20)"`
	ZhuokunTeacherCode string    `json:"zhuokun_teacher_code" xorm:"comment('卓坤医疗老师工号') VARCHAR(50)"`
	TeacherId          int64     `json:"teacher_id" xorm:"comment('华图在线老师ID') index BIGINT(20)"`
	UpdatedAt          time.Time `json:"updated_at" xorm:"DATETIME"`
	CreatedAt          time.Time `json:"created_at" xorm:"DATETIME"`
}
