package models

import (
	"time"
)

type TurnbackmoneyrecordLog struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Orderitemid int       `json:"OrderItemId" xorm:"not null INT(11)"`
	Remark      string    `json:"Remark" xorm:"not null VARCHAR(500)"`
	Optusername string    `json:"OptUserName" xorm:"not null VARCHAR(100)"`
	Optuserid   int       `json:"OptUserId" xorm:"not null INT(11)"`
	Addtime     time.Time `json:"AddTime" xorm:"not null DATETIME"`
}
