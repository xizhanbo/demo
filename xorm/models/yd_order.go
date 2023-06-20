package models

import (
	"time"
)

type YdOrder struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrderNum  string    `json:"order_num" xorm:"not null comment('在线订单编号') index VARCHAR(50)"`
	TrackId   string    `json:"track_id" xorm:"not null comment('云朵工单唯一id') unique VARCHAR(50)"`
	CreatedAt time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"comment('修改时间') DATETIME"`
}
