package models

import (
	"time"
)

type Collageorder struct {
	Id           int64     `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Groupid      int64     `json:"groupId" xorm:"not null comment('拼团ID') index BIGINT(20)"`
	Username     string    `json:"userName" xorm:"not null comment('团员用户名') VARCHAR(50)"`
	Userid       int       `json:"userId" xorm:"not null comment('团员ID') INT(11)"`
	Ordernum     string    `json:"orderNum" xorm:"not null comment('订单编号') index VARCHAR(50)"`
	Paystatus    int       `json:"payStatus" xorm:"not null default 1 comment('支付状态, 1-未支付, 2-已支付, 3-已退款, 4-取消支付5-支付失败 6退款中') index TINYINT(1)"`
	Createdat    time.Time `json:"createdAt" xorm:"not null comment('下单时间') DATETIME"`
	Payat        time.Time `json:"payAt" xorm:"comment('支付时间') DATETIME"`
	Statusat     time.Time `json:"statusAt" xorm:"comment('退款时间, 取消时间') DATETIME"`
	Money        int       `json:"money" xorm:"not null comment('订单金额, 分') INT(11)"`
	Activityid   int64     `json:"activityId" xorm:"not null comment('活动ID') index BIGINT(20)"`
	Address      string    `json:"address" xorm:"comment('详细地址') VARCHAR(200)"`
	Province     string    `json:"province" xorm:"comment('省份') VARCHAR(100)"`
	City         string    `json:"city" xorm:"comment('城市') VARCHAR(100)"`
	Area         string    `json:"area" xorm:"comment('地区') VARCHAR(100)"`
	Phone        string    `json:"phone" xorm:"comment('收件人手机号') CHAR(11)"`
	Consignee    string    `json:"consignee" xorm:"comment('收件人') VARCHAR(40)"`
	Isdelete     int       `json:"isDelete" xorm:"not null default 0 comment('0-未删除, 1-已删除') index TINYINT(1)"`
	Paymenttype  int       `json:"paymentType" xorm:"default 60 comment('60:小程序支付 32:支付宝支付 33:微信支付') TINYINT(1)"`
	Isrobotorder int       `json:"isRobotOrder" xorm:"not null default 0 comment('0-正常订单；1-机器人订单') TINYINT(1)"`
}
