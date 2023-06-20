package models

import (
	"time"
)

type Parnter struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ParnterName string    `json:"parnter_name" xorm:"not null comment('合作方名称') VARCHAR(50)"`
	OrderNumber string    `json:"order_number" xorm:"not null comment('订单编号头部') VARCHAR(10)"`
	Record      int       `json:"record" xorm:"not null default 0 comment('0 是  1否') TINYINT(2)"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('是否删除 0 否 1是') TINYINT(2)"`
	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(300)"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null DATETIME"`
	Creater     string    `json:"creater" xorm:"not null comment('创建人') VARCHAR(100)"`
}
