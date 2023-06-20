package models

import (
	"time"
)

type SonOrder struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ParentNum      string    `json:"parent_num" xorm:"not null comment('父订单号') index VARCHAR(100)"`
	SonNum         string    `json:"son_num" xorm:"not null comment('子订单号') index VARCHAR(100)"`
	Price          string    `json:"price" xorm:"not null comment('支付金额') DECIMAL(10,2)"`
	Status         int       `json:"status" xorm:"not null default 0 comment('0 未支付 1已支付') TINYINT(1)"`
	PayTime        time.Time `json:"pay_time" xorm:"comment('支付时间') DATETIME"`
	CreatedAt      time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"not null DATETIME"`
	RefundOrderNum string    `json:"refund_order_num" xorm:"comment('退款订单编号') VARCHAR(20)"`
}
