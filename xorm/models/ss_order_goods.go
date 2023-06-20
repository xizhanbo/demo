package models

import (
	"time"
)

type SsOrderGoods struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	OrderId   int       `json:"order_id" xorm:"not null comment('订单id') index INT(11)"`
	OrderNum  string    `json:"order_num" xorm:"not null comment('订单编号') index VARCHAR(200)"`
	GoodsId   int       `json:"goods_id" xorm:"not null comment('商品id') index INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"DATETIME"`
	ClassId   int       `json:"class_id" xorm:"default 0 comment('在线课id') INT(11)"`
}
