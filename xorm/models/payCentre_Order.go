package models

type PaycentreOrder struct {
	Id            int    `json:"id" xorm:"not null pk autoincr comment('主键') unique INT(10)"`
	Ordernum      string `json:"orderNum" xorm:"not null comment('支付时使用的业务订单编号') index VARCHAR(100)"`
	Requestid     string `json:"requestId" xorm:"not null comment('支付下单时标记每一笔支付订单的id，支付时随机生成') index VARCHAR(60)"`
	Outtradenum   string `json:"outTradeNum" xorm:"not null comment('支付中心支付单号') VARCHAR(30)"`
	Tradenum      string `json:"tradeNum" xorm:"not null comment('第三方支付单号，如支付宝') VARCHAR(30)"`
	Realpay       string `json:"realPay" xorm:"not null comment('支付金额') DECIMAL(10,2)"`
	Paytime       string `json:"payTime" xorm:"not null comment('支付时间，精确到毫秒的unix时间戳') CHAR(13)"`
	Childordernum string `json:"childOrderNum" xorm:"VARCHAR(50)"`
}
