package models

import (
	"time"
)

type Saleprogram struct {
	Rid          int       `json:"rid" xorm:"not null pk autoincr comment('自增活动ID') INT(11)"`
	Title        string    `json:"Title" xorm:"not null comment('活动名称') VARCHAR(200)"`
	Starttime    time.Time `json:"StartTime" xorm:"not null comment('开始时间') DATETIME"`
	Endtime      time.Time `json:"EndTime" xorm:"not null comment('结束时间') DATETIME"`
	Type         int       `json:"Type" xorm:"not null comment('活动类型1直降2打折4满减3组合优惠5优惠券') INT(11)"`
	Createtime   time.Time `json:"createtime" xorm:"not null comment('活动创建时间') DATETIME"`
	Optname      string    `json:"optname" xorm:"not null comment('操作人') VARCHAR(100)"`
	Optuserid    int       `json:"optUserId" xorm:"not null comment('操作用户ID') INT(11)"`
	Addstep1     int       `json:"addstep1" xorm:"comment('每隔一小时增加的购买量') INT(11)"`
	Addstep2     int       `json:"addstep2" xorm:"comment('每完成支付的订单增加的购买量') INT(11)"`
	Showtype     int       `json:"showType" xorm:"not null comment('显示类型0，每下一个订单 ，购买量加1    1：每隔一小时增加N个    2.每支付一个订单增加N个') INT(11)"`
	Urlwww       string    `json:"urlWWW" xorm:"comment('活动页地址（网站）') VARCHAR(500)"`
	Urlh5        string    `json:"urlH5" xorm:"comment('活动页地址（H5）') VARCHAR(500)"`
	Status       int       `json:"status" xorm:"not null comment('0，未上线，1已上线，2，下线,,4删除') INT(11)"`
	Code         string    `json:"code" xorm:"comment('活动编号') VARCHAR(50)"`
	Bz           string    `json:"bz" xorm:"comment('备注') VARCHAR(500)"`
	Validitytime time.Time `json:"ValidityTime" xorm:"comment('看课结束时间') DATETIME"`
	Validitydate int       `json:"ValidityDate" xorm:"comment('课程有效期') SMALLINT(5)"`
}
