package models

import (
	"time"
)

type ClOrderitem struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	Orderid       int64     `json:"orderId" xorm:"comment('订单id') BIGINT(20)"`
	Productid     int64     `json:"productId" xorm:"comment('课程id') BIGINT(20)"`
	Marketprice   int       `json:"marketPrice" xorm:"default 0 comment('市场价格') INT(11)"`
	Memberprice   int       `json:"memberPrice" xorm:"default 0 comment('单价价格') INT(11)"`
	Trueprice     int       `json:"truePrice" xorm:"default 0 comment('实际价格') INT(11)"`
	Buynum        int       `json:"buyNum" xorm:"default 0 comment('购买的数量') INT(11)"`
	Totalprice    int       `json:"totalPrice" xorm:"comment('付款价格') INT(11)"`
	Begintime     time.Time `json:"beginTime" xorm:"comment('下单时间') DATETIME"`
	Serviceterm   int       `json:"serviceTerm" xorm:"default 0 INT(11)"`
	Remark        string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(255)"`
	Presentexp    int       `json:"presentExp" xorm:"default 0 INT(255)"`
	Status        int       `json:"status" xorm:"comment('1:正常 2:退班 3:转班 4:删除 5:转班学员付款') INT(11)"`
	Statustime    time.Time `json:"statusTime" xorm:"comment('状态时间') DATETIME"`
	Iszengsong    int       `json:"isZengSong" xorm:"comment('1：随班赠送， 0：非赠送课程') TINYINT(3)"`
	Sid           int       `json:"sid" xorm:"comment('分校id') INT(11)"`
	Programid     int       `json:"programId" xorm:"comment('活动id') INT(11)"`
	Optuser       string    `json:"optUser" xorm:"default '' comment('操作人') VARCHAR(255)"`
	Turnbackmoney int       `json:"turnbackMoney" xorm:"comment('退款金额') INT(11)"`
	Ismakesure    int       `json:"isMakeSure" xorm:"comment('是否需要用户确认') TINYINT(3)"`
	Ismaizeng     int       `json:"isMaiZeng" xorm:"comment('是否含有买赠课程，1：含有 0：不含有') TINYINT(3)"`
}
