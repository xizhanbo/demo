package models

import (
	"time"
)

type ActivityWechatDelivery struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Title          string    `json:"title" xorm:"not null comment('活动标题') VARCHAR(64)"`
	YunduoTag      string    `json:"yunduo_tag" xorm:"not null comment('云朵标签') VARCHAR(32)"`
	WxappTitle     string    `json:"wxapp_title" xorm:"not null comment('小程序分享标题') VARCHAR(24)"`
	WxappImage     string    `json:"wxapp_image" xorm:"not null default '' comment('小程序分享封面') VARCHAR(192)"`
	Type           int       `json:"type" xorm:"not null default 1 comment('活动类型：1课程，2资料') TINYINT(4)"`
	Content        string    `json:"content" xorm:"not null default '' comment('资料链接或者课程id') VARCHAR(1024)"`
	CreatedAt      time.Time `json:"created_at" xorm:"not null DATETIME"`
	CreateOperator string    `json:"create_operator" xorm:"not null comment('创建人') VARCHAR(64)"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"not null DATETIME"`
	UpdateOperator string    `json:"update_operator" xorm:"not null comment('更新人') VARCHAR(64)"`
}
