package models

import (
	"time"
)

type Profit struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProxyId        int       `json:"proxy_id" xorm:"not null comment('收益人id	') index INT(11)"`
	OrderId        string    `json:"order_id" xorm:"not null comment('订单编号	') VARCHAR(200)"`
	MemberRegister time.Time `json:"member_register" xorm:"comment('学员注册时间') index DATETIME"`
	MemberName     string    `json:"member_name" xorm:"not null default '' comment('购买用户') VARCHAR(20)"`
	Price          string    `json:"price" xorm:"not null comment('实际支付价格') DECIMAL(10,2)"`
	PayTime        time.Time `json:"pay_time" xorm:"not null comment('支付时间') DATETIME"`
	Memo           string    `json:"memo" xorm:"comment('备注') TEXT"`
	Income         string    `json:"income" xorm:"default 0.00 comment('利润') DECIMAL(10,2)"`
	Status         int       `json:"status" xorm:"not null default 0 comment('收益状态	0待成熟 1可提现 2已申请提现 3已提现') index TINYINT(1)"`
	ParentId       int       `json:"parent_id" xorm:"not null default 0 comment('代理父级id') index INT(11)"`
	CreateTime     time.Time `json:"create_time" xorm:"not null comment('业绩时间') DATETIME"`
	IsPositive     int       `json:"is_positive" xorm:"not null default 0 comment('0 正向收益  1 负向收益') TINYINT(1)"`
	OrderStatus    int       `json:"order_status" xorm:"default 0 comment('0 正常   1退班  2转班 ') TINYINT(1)"`
	ClassId        int64     `json:"class_id" xorm:"not null comment('课程id') BIGINT(20)"`
	ClassTitle     string    `json:"class_title" xorm:"not null comment('课程标题') VARCHAR(100)"`
	ProxyStatus    int       `json:"proxy_status" xorm:"not null default 0 comment('代理收益状态0待成熟 1可提现 2已申请提现 3已提现') TINYINT(1)"`
}
