package models

import (
	"time"
)

type Studyrecord struct {
	Id           int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Classid      int       `json:"classId" xorm:"not null comment('课程ID') index INT(11)"`
	Lessonid     int       `json:"lessonId" xorm:"not null comment('课件ID') INT(11)"`
	Syllabusid   int       `json:"syllabusId" xorm:"not null comment('大纲ID') INT(11)"`
	Order        int       `json:"order" xorm:"not null default 0 comment('排序') INT(11)"`
	Username     string    `json:"userName" xorm:"not null comment('用户名') index(unionIndex) VARCHAR(50)"`
	Lasttime     time.Time `json:"lastTime" xorm:"not null default CURRENT_TIMESTAMP comment('最近学习时间') index index(unionIndex) TIMESTAMP"`
	Completetime int       `json:"completeTime" xorm:"not null comment('学习时长/ s') INT(11)"`
	Totaltime    int       `json:"totalTime" xorm:"not null comment('课件总时长/ s') INT(11)"`
	Type         int       `json:"type" xorm:"not null comment('1-直播,2-录播,3-回放') TINYINT(1)"`
	Origin       int       `json:"origin" xorm:"not null comment('1-华图在线, 2-面库') TINYINT(1)"`
	Isfinish     int       `json:"isFinish" xorm:"not null comment('0-未完成, 1-已完成') index TINYINT(1)"`
}
