package models

import (
	"time"
)

type Pushordertoedblog struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Ordernum   string    `json:"OrderNum" xorm:"not null comment('订单编号') index VARCHAR(255)"`
	Status     int       `json:"Status" xorm:"not null default 0 comment('未邮寄') TINYINT(4)"`
	Statusdate time.Time `json:"StatusDate" xorm:"not null default CURRENT_TIMESTAMP comment('状态更新时间') TIMESTAMP"`
	Orderid    int       `json:"OrderId" xorm:"not null comment('订单ID') index INT(11)"`
	Createdate time.Time `json:"CreateDate" xorm:"not null DATETIME"`
	Iscool     int       `json:"isCool" xorm:"not null default 0 comment('是否是面酷') TINYINT(4)"`
}
