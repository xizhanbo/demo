package models

import (
	"time"
)

type MemberRecordLog struct {
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	UserName        string    `json:"user_name" xorm:"not null default '' comment('用户名') index VARCHAR(50)"`
	MemberStartTime time.Time `json:"member_start_time" xorm:"not null default '0000-00-00 00:00:00' comment('会员开始时间') DATETIME"`
	MemberEndTime   time.Time `json:"member_end_time" xorm:"not null default '0000-00-00 00:00:00' comment('会员结束时间') DATETIME"`
}
