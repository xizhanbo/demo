package models

import (
	"time"
)

type MsgReceivebox struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Msgid    int       `json:"msgid" xorm:"not null INT(10)"`
	Userid   int       `json:"userid" xorm:"not null index INT(10)"`
	Isread   int       `json:"isread" xorm:"not null default 0 comment('0:未读 1:已读') TINYINT(1)"`
	Readdate time.Time `json:"readdate" xorm:"not null DATETIME"`
	Status   int       `json:"status" xorm:"not null default 1 comment('1:正常 3:删除') TINYINT(1)"`
}
