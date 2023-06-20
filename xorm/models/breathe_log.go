package models

import (
	"time"
)

type BreatheLog struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	CallUuid    string    `json:"call_uuid" xorm:"not null comment('会话唯一标识') VARCHAR(100)"`
	Caller      string    `json:"caller" xorm:"comment('主叫号码') VARCHAR(50)"`
	Callee      string    `json:"callee" xorm:"comment('手机号') CHAR(11)"`
	Direction   int       `json:"direction" xorm:"not null default 2 comment('1:呼入 2 呼出') index TINYINT(1)"`
	TaskResult  string    `json:"task_result" xorm:"comment('外呼结果') VARCHAR(50)"`
	TagName     string    `json:"tag_name" xorm:"comment('意向') VARCHAR(100)"`
	TagMax      int       `json:"tag_max" xorm:"comment('最高意向度') index TINYINT(1)"`
	IsPush      int       `json:"is_push" xorm:"not null default 0 comment('0 未推送  1 已推送') index TINYINT(1)"`
	SensorsPush int       `json:"sensors_push" xorm:"not null default 0 comment('0 未推送  1 已推送') TINYINT(1)"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
