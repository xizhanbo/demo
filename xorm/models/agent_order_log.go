package models

import (
	"time"
)

type AgentOrderLog struct {
	Id        int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(20)"`
	AgentName string    `json:"agent_name" xorm:"not null comment('代理商用户名') index VARCHAR(50)"`
	OrderId   int64     `json:"order_id" xorm:"not null comment('代报订单ID') BIGINT(20)"`
	ClassId   int       `json:"class_id" xorm:"not null comment('课程ID') INT(11)"`
	PayType   int       `json:"pay_type" xorm:"not null default 0 comment('支付方式：0学习卡，1虚拟币') INT(11)"`
	Remark    string    `json:"remark" xorm:"not null default '' comment('备注') VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at" xorm:"comment('下单时间') DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"comment('更新时间') DATETIME"`
}
