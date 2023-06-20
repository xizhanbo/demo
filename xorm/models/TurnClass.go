package models

type Turnclass struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Orderid       int64  `json:"orderid" xorm:"comment('老班级id') index BIGINT(20)"`
	Newordernum   string `json:"newOrderNum" xorm:"default '' comment('新班级订单号') index VARCHAR(50)"`
	Bookdeduction string `json:"bookDeduction" xorm:"default 0.00 comment('图书课时扣费') DECIMAL(10,2)"`
	Newclassprice string `json:"newClassPrice" xorm:"default 0.00 comment('新班价格') DECIMAL(10,2)"`
	Oldorderprice string `json:"oldOrderPrice" xorm:"default 0.00 comment('老订单价格') DECIMAL(10,2)"`
}
