package models

import (
	"time"
)

type CourseCard struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Title         string    `json:"title" xorm:"not null comment('课程卡名称') VARCHAR(150)"`
	GivingCourses string    `json:"giving_courses" xorm:"comment('赠送课程（多个英文逗号隔开）') TEXT"`
	CreateTime    time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	Status        int       `json:"status" xorm:"default 1 comment('状态：1正常  2:删除') index TINYINT(3)"`
	Creator       string    `json:"creator" xorm:"comment('创建人') VARCHAR(50)"`
}
