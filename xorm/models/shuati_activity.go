package models

import (
	"time"
)

type ShuatiActivity struct {
	Id             int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ClassId        int64     `json:"class_id" xorm:"not null comment('刷题班活动绑定的课程id') index BIGINT(20)"`
	PunchStartTime time.Time `json:"punch_start_time" xorm:"not null comment('刷题班活动开始时间') DATETIME"`
	PunchEndTime   time.Time `json:"punch_end_time" xorm:"not null comment('刷题班活动打卡结束时间') DATETIME"`
	EndTime        time.Time `json:"end_time" xorm:"not null comment('刷题班活动结束时间') DATETIME"`
	PunchDays      int       `json:"punch_days" xorm:"not null default 0 comment('返现最低打卡天数') SMALLINT(6)"`
	Status         string    `json:"status" xorm:"not null comment('活动状态: draft未上线，publish上线，offline已下线') VARCHAR(16)"`
	CreatedAt      time.Time `json:"created_at" xorm:"TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"TIMESTAMP"`
	Operator       string    `json:"operator" xorm:"default '' comment('操作人UserName') VARCHAR(64)"`
}
