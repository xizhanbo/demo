package models

import (
	"time"
)

type ZhuokunCourseForOnlineCourses struct {
	Id              int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ZhuokunCourseId int64     `json:"zhuokun_course_id" xorm:"comment('卓坤课程ID') index BIGINT(20)"`
	OnlineCourseId  int64     `json:"online_course_id" xorm:"comment('华图在线课程ID') index BIGINT(20)"`
	Type            int       `json:"type" xorm:"comment('1非套餐课 2套餐课') TINYINT(3)"`
	ExpireType      string    `json:"expire_type" xorm:"comment('学习有效期类型') ENUM('day','end','fixed','forever')"`
	ExpireStart     time.Time `json:"expire_start" xorm:"comment('有效期开始') DATETIME"`
	ExpireEnd       time.Time `json:"expire_end" xorm:"comment('有效期结束') DATETIME"`
	ExpireDays      int       `json:"expire_days" xorm:"comment('有效天数 整型') INT(10)"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"DATETIME"`
	CreatedAt       time.Time `json:"created_at" xorm:"DATETIME"`
}
