package models

import (
	"time"
)

type NewMessage struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type          int       `json:"type" xorm:"not null comment('消息类型') unique TINYINT(1)"`
	CreateAt      string    `json:"create_at" xorm:"DECIMAL(10)"`
	UpdateAt      time.Time `json:"update_at" xorm:"DATETIME"`
	CreatorName   string    `json:"creator_name" xorm:"not null comment('操作人') VARCHAR(100)"`
	MessageConent string    `json:"message_conent" xorm:"comment('消息内容') TEXT"`
}
