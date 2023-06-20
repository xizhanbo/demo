package models

import (
	"time"
)

type Redenvelopeactivity struct {
	Id               int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	Activename       string    `json:"activeName" xorm:"default '' comment('活动名称') VARCHAR(100)"`
	Createtime       time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	Begintime        time.Time `json:"beginTime" xorm:"comment('活动开始时间') DATETIME"`
	Endtime          time.Time `json:"endTime" xorm:"comment('活动结束时间') DATETIME"`
	Activelabel      string    `json:"activeLabel" xorm:"default '' comment('活动标签') VARCHAR(50)"`
	Labledesc        string    `json:"lableDesc" xorm:"default '' comment('标签描述') VARCHAR(100)"`
	Alonebyprice     int       `json:"aloneByPrice" xorm:"comment('单独购买金额，单位分') INT(11)"`
	Alonenum         int       `json:"aloneNum" xorm:"comment('单独购买个数') INT(11)"`
	Aloneminiprice   int       `json:"aloneMiniPrice" xorm:"comment('单独购买红包最小金额，单位分') INT(11)"`
	Collagebyprice   int       `json:"collageByPrice" xorm:"comment('拼团红包金额') INT(11)"`
	Collagenum       int       `json:"collageNum" xorm:"comment('拼团红包个数') INT(11)"`
	Collageminiprice int       `json:"collageMiniPrice" xorm:"comment('拼团红包最少金额，单位：分') INT(11)"`
	Status           int       `json:"status" xorm:"default 0 comment('0:未开始，1开始，2失效，3结束，4删除') TINYINT(3)"`
	Classids         string    `json:"classIds" xorm:"TEXT"`
}
