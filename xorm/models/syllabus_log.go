package models

import (
	"time"
)

type SyllabusLog struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserName   string    `json:"user_name" xorm:"not null comment('用户名') index VARCHAR(100)"`
	UserId     int       `json:"user_id" xorm:"not null comment('用户id') index INT(11)"`
	NetClassId int       `json:"net_class_id" xorm:"not null comment('课程id') index INT(11)"`
	Type       int       `json:"type" xorm:"default 0 comment('1添加2修改3删除') TINYINT(1)"`
	Content    string    `json:"content" xorm:"comment('操作内容') TEXT"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DATETIME"`
}
