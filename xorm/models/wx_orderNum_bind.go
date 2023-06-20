package models

type WxOrdernumBind struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	OutTradeNo string `json:"out_trade_no" xorm:"not null default '' comment('预约号') unique VARCHAR(50)"`
	Ordernum   string `json:"orderNum" xorm:"not null default '' comment('订单编号') unique VARCHAR(50)"`
}
