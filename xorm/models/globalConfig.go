package models

import (
	"time"
)

type Globalconfig struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Title     string    `json:"title" xorm:"not null comment('配置标题') VARCHAR(40)"`
	Key       string    `json:"key" xorm:"not null comment('配置建名') unique VARCHAR(100)"`
	Value     string    `json:"value" xorm:"not null comment('配置值') VARCHAR(200)"`
	Cate      int       `json:"cate" xorm:"not null comment('配置分类, 1-课程, 2-搜索,3-图书') TINYINT(4)"`
	Timestamp time.Time `json:"timestamp" xorm:"not null default CURRENT_TIMESTAMP comment('时间戳,自动更新') TIMESTAMP"`
	Type      string    `json:"type" xorm:"default '' VARCHAR(20)"`
	Catename  string    `json:"cateName" xorm:"comment('分类名称') VARCHAR(100)"`
}
