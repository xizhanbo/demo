package models

import (
	"time"
)

type ShuatiUserWithdraw struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId     int64     `json:"user_id" xorm:"not null comment('用户id') index BIGINT(20)"`
	ClassId    int64     `json:"class_id" xorm:"not null comment('用户id') index BIGINT(20)"`
	Fee        string    `json:"fee" xorm:"not null comment('返现金额') DECIMAL(8,2)"`
	Billno     string    `json:"billno" xorm:"not null comment('流水号') VARCHAR(64)"`
	Status     int       `json:"status" xorm:"not null comment('状态：1成功，2失败') TINYINT(4)"`
	StatusCode string    `json:"status_code" xorm:"not null comment('微信转账接口返回状态码') VARCHAR(32)"`
	CreatedAt  time.Time `json:"created_at" xorm:"TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"TIMESTAMP"`
}
