package models

import (
	"time"
)

type Thirdorderbind struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	Orderid       int64     `json:"orderId" xorm:"comment('在线订单id') BIGINT(20)"`
	Ordernum      string    `json:"orderNum" xorm:"comment('在线订单编号') index VARCHAR(50)"`
	Thirdordernum string    `json:"thirdOrderNum" xorm:"comment('第三方订单编号') index VARCHAR(50)"`
	Orderorigin   int       `json:"orderOrigin" xorm:"comment('订单来源') TINYINT(3)"`
	Addtime       time.Time `json:"addTime" xorm:"comment('添加时间') TIMESTAMP"`
}
