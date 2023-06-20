package models

import (
	"time"
)

type CoursewareStudyRecords struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId        int       `json:"user_id" xorm:"comment('Java用户中心的UserId') index INT(11)"`
	Username      string    `json:"username" xorm:"not null comment('Java用户中心的UserName') unique(user_lesson) VARCHAR(50)"`
	CoursewareId  int       `json:"courseware_id" xorm:"not null comment('课件ID') unique(user_lesson) INT(11)"`
	Type          int       `json:"type" xorm:"not null comment('1-直播,2-录播,3-回放') TINYINT(1)"`
	Origin        int       `json:"origin" xorm:"not null comment('1-华图在线, 2-面库') TINYINT(1)"`
	CompletedTime int       `json:"completed_time" xorm:"not null comment('学习时长/ s') INT(11)"`
	TotalTime     int       `json:"total_time" xorm:"not null comment('课件总时长/ s') INT(11)"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"not null comment('最近学习时间') index DATETIME"`
	IsFinished    int       `json:"is_finished" xorm:"not null comment('0-未完成, 1-已完成') TINYINT(1)"`
}
