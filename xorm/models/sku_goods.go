package models

import (
	"time"
)

type SkuGoods struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CollectionId int       `json:"collection_id" xorm:"not null default 0 comment('合集id') index INT(11)"`
	ClassId      int       `json:"class_id" xorm:"not null default 0 comment('课程id') INT(11)"`
	SkuPath      string    `json:"sku_path" xorm:"not null comment('子集路径') VARCHAR(255)"`
	CreatedAt    time.Time `json:"created_at" xorm:"not null comment('添加时间') DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"not null comment('修改时间') DATETIME"`
	Creator      string    `json:"creator" xorm:"not null comment('添加人') VARCHAR(100)"`
}
