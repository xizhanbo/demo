package models

import (
	"time"
)

type SrmUserExt struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserName     string    `json:"user_name" xorm:"default '' comment('用户名') index VARCHAR(100)"`
	Mobile       int64     `json:"mobile" xorm:"not null comment('手机号码') index BIGINT(11)"`
	Operator     string    `json:"operator" xorm:"not null comment('操作人') VARCHAR(100)"`
	WechatId     string    `json:"wechat_id" xorm:"default '' comment('微信id') VARCHAR(50)"`
	WechatNum    string    `json:"wechat_num" xorm:"comment('微信号') VARCHAR(50)"`
	WechatOpenid string    `json:"wechat_openid" xorm:"comment('微信openid') VARCHAR(160)"`
	LastChoice   int       `json:"last_choice" xorm:"default 0 comment('是否选中0否1是') TINYINT(1)"`
	CreatedAt    time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"DATETIME"`
}
