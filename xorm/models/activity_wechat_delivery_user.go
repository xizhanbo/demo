package models

import (
	"time"
)

type ActivityWechatDeliveryUser struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ActivityId int       `json:"activity_id" xorm:"not null comment('活动id') INT(11)"`
	UserId     int64     `json:"user_id" xorm:"not null comment('用户id') index BIGINT(20)"`
	Mobile     string    `json:"mobile" xorm:"not null comment('手机号') VARCHAR(16)"`
	Unionid    string    `json:"unionid" xorm:"not null comment('微信unionid') VARCHAR(64)"`
	CreatedAt  time.Time `json:"created_at" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}
