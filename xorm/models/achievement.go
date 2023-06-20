package models

import (
	"time"
)

type Achievement struct {
	Id             int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Renshu         int       `json:"renshu" xorm:"not null default 0 comment('招生人数（当日已到账的月卡订单数+普通订单数）') INT(11)"`
	Yeji           string    `json:"yeji" xorm:"not null default 0 comment('业绩（当日所有业绩）') DECIMAL(10)"`
	Cashorders     int       `json:"cashOrders" xorm:"not null default 0 comment('使用现金支付的订单(人)数') INT(11)"`
	Monthcardorder int       `json:"monthCardOrder" xorm:"not null default 0 comment('月卡订单数') INT(11)"`
	Riqi           time.Time `json:"riqi" xorm:"not null comment('日期') DATE"`
	Status         int       `json:"status" xorm:"not null default 0 comment('状态') TINYINT(4)"`
}
