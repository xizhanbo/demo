package models

import (
	"time"
)

type Collageactivity struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	Classid        int       `json:"classId" xorm:"comment('课程id') index INT(11)"`
	Createat       time.Time `json:"createAt" xorm:"comment('创建时间') DATETIME"`
	Beginat        time.Time `json:"beginAt" xorm:"comment('活动开始时间') DATETIME"`
	Endat          time.Time `json:"endAt" xorm:"comment('活动结束时间') index DATETIME"`
	Collagepeople  int       `json:"collagePeople" xorm:"comment('拼团人数') SMALLINT(5)"`
	Allowcollage   int       `json:"allowCollage" xorm:"comment('前端展示加入团开关：0(不允许直接拼团) 1(允许直接拼团)') TINYINT(1)"`
	Limittype      int       `json:"limitType" xorm:"default 0 comment('参团限制0:不限制,1:只允许新人参加,2:只允许老学员参加') TINYINT(1)"`
	Discountprice  int       `json:"discountPrice" xorm:"comment('优惠价格（单位：分）') INT(11)"`
	Collageprice   int       `json:"collagePrice" xorm:"comment('拼团价格（单位：分）') INT(11)"`
	Showtype       int       `json:"showType" xorm:"default 0 comment('已购人数设置:0:实际购买人数,1:每小时添加,2:购买倍数添加') TINYINT(1)"`
	Status         int       `json:"status" xorm:"default 0 comment('活动状态:0:未上线,1:已上线,2:已失效,3:已结束') index TINYINT(1)"`
	Openname       string    `json:"openName" xorm:"default '' comment('创建人') VARCHAR(50)"`
	Increasepeople int       `json:"increasePeople" xorm:"default 0 comment('增长人数') INT(11)"`
	Activitylength int       `json:"activityLength" xorm:"comment('拼团有效时长') INT(11)"`
	Classtitle     string    `json:"classTitle" xorm:"comment('课程标题') VARCHAR(100)"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"comment('修改时间') DATETIME"`
	Robotstatus    int       `json:"robotStatus" xorm:"default 0 comment('机器人开关 0:关 1:开') TINYINT(1)"`
	Startlimittype int       `json:"startLimitType" xorm:"default 0 comment('开团人限制：0:不限制,1:只允许新人参加,2:只允许老学员参加') TINYINT(1)"`
}
