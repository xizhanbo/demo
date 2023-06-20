package models

import (
	"time"
)

type CoffeeOrder struct {
	Id           int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(11)"`
	UserName     string    `json:"user_name" xorm:"not null default '' comment('用户名') index VARCHAR(50)"`
	OrderId      int64     `json:"order_id" xorm:"not null comment('FK:Cl_Order主键ID') index BIGINT(11)"`
	PayType      int       `json:"pay_type" xorm:"not null default 0 comment('支付方式: 支付宝32 微信33') TINYINT(1)"`
	PayInAdvance string    `json:"pay_in_advance" xorm:"not null default 0.00 comment('首付款') DECIMAL(12,2)"`
	ApplyCode    string    `json:"apply_code" xorm:"not null default '' comment('咖啡易融申请码') index CHAR(6)"`
	IsDestroy    int       `json:"is_destroy" xorm:"not null default 0 comment('申请码是否核验 0否 1是') TINYINT(1)"`
	CreateTime   time.Time `json:"create_time" xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `json:"update_time" xorm:"default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}
