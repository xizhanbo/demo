package models

import (
	"time"
)

type Redenvelopeputforward struct {
	Id              int64     `json:"id" xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	Putforwardtime  time.Time `json:"putForwardTime" xorm:"comment('提现时间') DATETIME"`
	Putforwardprice int       `json:"putForwardPrice" xorm:"comment('提现金额') INT(11)"`
	Username        string    `json:"userName" xorm:"default '' comment('用户名') VARCHAR(50)"`
	Status          int       `json:"status" xorm:"default 1 comment('状态') TINYINT(3)"`
	Putforwardtype  int       `json:"putForwardType" xorm:"comment('提现支付方式, 1-微信, 2-支付宝') TINYINT(2)"`
	Mchid           string    `json:"mchId" xorm:"default '' comment('提现商户ID') VARCHAR(30)"`
	Ordernum        string    `json:"orderNum" xorm:"default '' comment('提现订单ID') VARCHAR(30)"`
	Account         string    `json:"account" xorm:"default '' comment('提现账号, 微信记录openid, 支付宝记录支付宝账号') VARCHAR(50)"`
	Mobile          string    `json:"mobile" xorm:"comment('提现手机号') CHAR(11)"`
}
