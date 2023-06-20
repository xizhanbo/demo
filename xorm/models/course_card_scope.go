package models

type CourseCardScope struct {
	Id     int    `json:"id" xorm:"not null pk autoincr comment('课程卡适用范围ID') INT(10)"`
	Name   string `json:"name" xorm:"comment('课程卡适用范围名称') VARCHAR(50)"`
	Status int    `json:"status" xorm:"default 1 comment('状态') TINYINT(3)"`
}
