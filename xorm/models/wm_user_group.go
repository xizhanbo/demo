package models

import (
	"time"
)

type WmUserGroup struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	TaskId       int       `json:"task_id" xorm:"not null comment('用户分群任务id') index INT(11)"`
	UserName     string    `json:"user_name" xorm:"not null comment('用户名') index VARCHAR(30)"`
	Phone        string    `json:"phone" xorm:"not null comment('用户手机号') index VARCHAR(20)"`
	WechatNum    string    `json:"wechat_num" xorm:"comment('用户微信号') VARCHAR(50)"`
	WechatOpenid string    `json:"wechat_openid" xorm:"default '' comment('用户微信openId') VARCHAR(160)"`
	WechatId     string    `json:"wechat_id" xorm:"comment('用户微信Id') VARCHAR(50)"`
	CreatedAt    time.Time `json:"created_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
}
