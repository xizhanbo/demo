package models

import (
	"time"
)

type Redenvelope struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	Activityid     int64     `json:"activityId" xorm:"comment('活动id') BIGINT(20)"`
	Sendname       string    `json:"sendName" xorm:"default '' comment('发红包人') VARCHAR(50)"`
	Endtime        time.Time `json:"endTime" xorm:"comment('红包过期时间') DATETIME"`
	Type           int       `json:"type" xorm:"default 0 comment('红包类型：0普通购买，1拼团') TINYINT(3)"`
	Classid        int64     `json:"classId" xorm:"comment('课程id') BIGINT(20)"`
	Sumnum         int       `json:"sumNum" xorm:"comment('红包总数') INT(11)"`
	Surplusnum     int       `json:"surplusNum" xorm:"comment('剩余数量') INT(11)"`
	Sumprice       int       `json:"sumPrice" xorm:"comment('红包总金额') INT(11)"`
	Surplusprice   int       `json:"surplusPrice" xorm:"comment('剩余总金额') INT(11)"`
	Aloneminiprice int       `json:"aloneMiniPrice" xorm:"comment('单个红包最小金额') INT(11)"`
	Orderid        int       `json:"orderId" xorm:"comment('订单id') INT(11)"`
	Status         int       `json:"status" xorm:"default 0 comment('状态，1有效，0无效，2：订单已经退班, 3-订单已经转班') TINYINT(3)"`
	Classtitle     string    `json:"classTitle" xorm:"default '' comment('课程标题') VARCHAR(100)"`
}
