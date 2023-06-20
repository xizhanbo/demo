package models

import (
	"time"
)

type Withdrawals struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProxyId        int       `json:"proxy_id" xorm:"not null comment('代理人id') index INT(11)"`
	UserName       string    `json:"user_name" xorm:"not null comment('姓名') VARCHAR(20)"`
	UserMobile     string    `json:"user_mobile" xorm:"not null comment('手机号') VARCHAR(11)"`
	Price          string    `json:"price" xorm:"default 0.00 comment('到账金额	') DECIMAL(10,2)"`
	PriceWhere     string    `json:"price_where" xorm:"not null comment('资金去向') VARCHAR(100)"`
	DirectSale     string    `json:"direct_sale" xorm:"not null comment('个人直售') DECIMAL(10,2)"`
	AllySale       string    `json:"ally_sale" xorm:"not null comment('盟友贡献') DECIMAL(10,2)"`
	DirectSaleBack string    `json:"direct_sale_back" xorm:"not null comment('个人直售转退班') DECIMAL(10,2)"`
	AllySaleBack   string    `json:"ally_sale_back" xorm:"not null comment('盟友贡献转退班') DECIMAL(10,2)"`
	Tax            string    `json:"tax" xorm:"not null comment('个税') DECIMAL(10,2)"`
	Achievement    string    `json:"achievement" xorm:"not null comment('业绩') DECIMAL(10,2)"`
	PayTime        time.Time `json:"pay_time" xorm:"comment('返款时间') DATETIME"`
	AuditStatus    int       `json:"audit_status" xorm:"not null default 0 comment('提现状态 0待审核 1审核未通过 2审核已通过 3打款成功 ') index TINYINT(1)"`
	ApplyTime      time.Time `json:"apply_time" xorm:"not null comment('申请时间') DATETIME"`
	AuditName      string    `json:"audit_name" xorm:"comment('操作人') VARCHAR(20)"`
	Operationdate  time.Time `json:"operationDate" xorm:"comment('操作时间') DATETIME"`
}
