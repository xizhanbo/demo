package models

import (
	"time"
)

type SkuCollection struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title        string    `json:"title" xorm:"not null comment('名称') VARCHAR(100)"`
	CollectionId int       `json:"collection_id" xorm:"not null default 0 comment('合集id') index INT(11)"`
	Pid          int       `json:"pid" xorm:"not null default 0 comment('父级id') INT(11)"`
	CreatedAt    time.Time `json:"created_at" xorm:"not null comment('添加时间') DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"not null comment('修改时间') DATETIME"`
	Creator      string    `json:"creator" xorm:"not null comment('操作人') VARCHAR(100)"`
}
