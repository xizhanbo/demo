package models

import (
	"time"
)

type YouzanAutosyncOrder struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Msg            string    `json:"msg" xorm:"not null VARCHAR(255)"`
	YouzanOrdernum string    `json:"youzan_ordernum" xorm:"not null index VARCHAR(50)"`
	PushTime       time.Time `json:"push_time" xorm:"not null index DATETIME"`
	Phone          string    `json:"phone" xorm:"not null index CHAR(11)"`
	Status         int       `json:"status" xorm:"default 0 comment('是否查看（0）待处理，（1）已处理') TINYINT(1)"`
	Operator       string    `json:"operator" xorm:"default '' comment('操作人') VARCHAR(50)"`
	UpdateAt       time.Time `json:"update_at" xorm:"comment('操作时间') DATETIME"`
}
