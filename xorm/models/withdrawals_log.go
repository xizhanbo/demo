package models

import (
	"time"
)

type WithdrawalsLog struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	WithId         int       `json:"with_id" xorm:"not null comment('审核结算Id') index INT(10)"`
	UserName       string    `json:"user_name" xorm:"not null comment('代理人') VARCHAR(20)"`
	DirectSale     string    `json:"direct_sale" xorm:"not null comment('个人直售价格') DECIMAL(10,2)"`
	AllySale       string    `json:"ally_sale" xorm:"not null comment('盟友贡献价格') DECIMAL(10,2)"`
	DirectSaleBack string    `json:"direct_sale_back" xorm:"not null comment('个人直售转退班价格') DECIMAL(10,2)"`
	AllySaleBack   string    `json:"ally_sale_back" xorm:"not null comment('盟友直售转退班价格') DECIMAL(10,2)"`
	Tax            string    `json:"tax" xorm:"not null comment('个税') DECIMAL(10,2)"`
	AuditName      string    `json:"audit_name" xorm:"not null comment('操作人姓名') VARCHAR(20)"`
	Remarks        string    `json:"remarks" xorm:"comment('备注') TEXT"`
	Status         int       `json:"status" xorm:"not null comment('操作状态') index TINYINT(1)"`
	Operationdate  time.Time `json:"operationDate" xorm:"not null comment('操作时间') DATETIME"`
	Achievement    string    `json:"achievement" xorm:"not null comment('业绩') DECIMAL(10,2)"`
	PriceWhere     string    `json:"price_where" xorm:"not null comment('资金去向') VARCHAR(100)"`
}
