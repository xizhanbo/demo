package models

import (
	"time"
)

type ProtocolTemplate struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProtocolNum  int64     `json:"protocol_num" xorm:"not null comment('契约锁id') index BIGINT(20)"`
	ProtocolName string    `json:"protocol_name" xorm:"not null comment('协议模板名称') index VARCHAR(200)"`
	Status       int       `json:"status" xorm:"not null comment('状态 1启用0禁用2删除') TINYINT(4)"`
	CreateAt     time.Time `json:"create_at" xorm:"DATETIME"`
	UpdateAt     time.Time `json:"update_at" xorm:"DATETIME"`
	CreateName   string    `json:"create_name" xorm:"not null default '' comment('创建人') VARCHAR(200)"`
}
