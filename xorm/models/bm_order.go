package models

type BmOrder struct {
	Id      int    `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Orderid int    `json:"orderId" xorm:"not null comment('订单id') index INT(11)"`
	City    string `json:"city" xorm:"default '' comment('地区') index VARCHAR(255)"`
	Fx      string `json:"fx" xorm:"default '' comment('分校') index VARCHAR(255)"`
	Source  string `json:"source" xorm:"default '' comment('来源（bm：教育）') index VARCHAR(100)"`
}
