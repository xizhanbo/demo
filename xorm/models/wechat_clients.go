package models

import (
	"time"
)

type WechatClients struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	WechatUserId     int       `json:"wechat_user_id" xorm:"not null index INT(11)"`
	OpenId           string    `json:"open_id" xorm:"not null unique(open_id) index VARCHAR(191)"`
	Type             int       `json:"type" xorm:"comment('1wechat(公众号)2miniprogram(小程序)') unique(open_id) TINYINT(1)"`
	SessionKey       string    `json:"session_key" xorm:"index VARCHAR(191)"`
	Subscribe        int       `json:"subscribe" xorm:"not null default 0 comment('0未关注1已关注') TINYINT(1)"`
	SubscribeAt      time.Time `json:"subscribe_at" xorm:"comment('关注时间') DATETIME"`
	UnsubscribeAt    time.Time `json:"unsubscribe_at" xorm:"comment('取消关注时间') DATETIME"`
	FirstSubscribeAt time.Time `json:"first_subscribe_at" xorm:"comment('首次关注时间') DATETIME"`
	CreatedAt        time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"DATETIME"`
}
