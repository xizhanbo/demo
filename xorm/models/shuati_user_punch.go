package models

import (
	"time"
)

type ShuatiUserPunch struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ClassId   int64     `json:"class_id" xorm:"not null comment('刷题活动课程id') index BIGINT(20)"`
	UserId    int64     `json:"user_id" xorm:"not null comment('用户id') index BIGINT(20)"`
	PunchDay  time.Time `json:"punch_day" xorm:"not null comment('打卡日期') index DATE"`
	Channel   int       `json:"channel" xorm:"not null default 1 comment('打卡渠道：1小打卡') TINYINT(4)"`
	Operator  string    `json:"operator" xorm:"not null default '' comment('操作员') VARCHAR(64)"`
	Status    int       `json:"status" xorm:"not null default 1 comment('打卡状态：1成功，2未打卡') TINYINT(4)"`
	CreatedAt time.Time `json:"created_at" xorm:"TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"TIMESTAMP"`
}
