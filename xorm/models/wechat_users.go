package models

import (
	"time"
)

type WechatUsers struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(10)"`
	UnionId   string    `json:"union_id" xorm:"not null comment('用户唯一id') unique index VARCHAR(150)"`
	Nickname  string    `json:"nickname" xorm:"comment('用户昵称') VARCHAR(150)"`
	Sex       int       `json:"sex" xorm:"comment('性别0未知1男2女') TINYINT(1)"`
	Province  string    `json:"province" xorm:"comment('省份') VARCHAR(150)"`
	City      string    `json:"city" xorm:"comment('城市') VARCHAR(150)"`
	Country   string    `json:"country" xorm:"comment('国家') VARCHAR(150)"`
	Address   string    `json:"address" xorm:"comment('街道') VARCHAR(191)"`
	Avatar    string    `json:"avatar" xorm:"comment('头像url') index VARCHAR(191)"`
	CreatedAt time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"DATETIME"`
}
