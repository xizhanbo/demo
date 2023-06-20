package models

import (
	"time"
)

type CourseCollection struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('递增主键') INT(11)"`
	Username   string    `json:"userName" xorm:"not null comment('用户名') index VARCHAR(255)"`
	Netclassid int       `json:"netClassId" xorm:"not null comment('收藏课程ID') index INT(11)"`
	Addtime    time.Time `json:"addTime" xorm:"comment('收藏时间') DATETIME"`
	Terminal   int       `json:"terminal" xorm:"default 1 comment('1.android;2.ios;3.pc;4.ipad android;5.ipad ios; 6.weixin 7m站21微信小程序') TINYINT(4)"`
}
