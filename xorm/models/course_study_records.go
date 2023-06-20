package models

import (
	"time"
)

type CourseStudyRecords struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	UserId    int       `json:"user_id" xorm:"comment('Java用户中心的UserId') index INT(11)"`
	Username  string    `json:"username" xorm:"not null comment('Java用户中心的UserName') unique(username_class) VARCHAR(50)"`
	CourseId  int       `json:"course_id" xorm:"not null default 0 comment('NetClasses表rid') unique(username_class) INT(11)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') index DATETIME"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
}
