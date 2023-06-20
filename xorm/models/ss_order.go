package models

import (
	"time"
)

type SsOrder struct {
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OrderNum       string    `json:"order_num" xorm:"not null comment('业务订单编号') index VARCHAR(100)"`
	HhrOrderNum    string    `json:"hhr_order_num" xorm:"default '' comment('合伙人订单号') VARCHAR(100)"`
	TurnOrderNum   string    `json:"turn_order_num" xorm:"default '' comment('转班新订单编号') index VARCHAR(100)"`
	ReturnOrderNum string    `json:"return_order_num" xorm:"default '' comment('退班订单编号') VARCHAR(100)"`
	SsOrderId      int       `json:"ss_order_id" xorm:"comment('神一订单ID') INT(11)"`
	SsOrderNum     string    `json:"ss_order_num" xorm:"default '' comment('神一订单编号') VARCHAR(100)"`
	SsRefundNum    string    `json:"ss_refund_num" xorm:"default '' comment('神一退单编号') VARCHAR(100)"`
	UpdateTime     time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
}
