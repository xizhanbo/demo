package models

import (
	"time"
)

type NetlessionsLog struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Userid     int64     `json:"userId" xorm:"not null comment('用户id') BIGINT(20)"`
	Username   string    `json:"userName" xorm:"not null comment('实名') index VARCHAR(200)"`
	Netclassid int64     `json:"NetClassId" xorm:"not null comment('课程id') BIGINT(20)"`
	Netid      string    `json:"NetId" xorm:"not null comment('课件id') VARCHAR(255)"`
	Content    string    `json:"content" xorm:"not null TEXT"`
	Tag        string    `json:"tag" xorm:"not null VARCHAR(255)"`
	Logdate    time.Time `json:"logDate" xorm:"not null DATETIME"`
	Logtime    int       `json:"logTime" xorm:"not null INT(11)"`
}
