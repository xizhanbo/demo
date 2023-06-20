package models

import (
	"time"
)

type Redenvelopereceive struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	Openname      string    `json:"openName" xorm:"default '' comment('领取者用户') VARCHAR(50)"`
	Sendname      string    `json:"sendName" xorm:"default '' comment('发送者用户') VARCHAR(50)"`
	Redenvelopeid int       `json:"redEnvelopeId" xorm:"comment('红包发送表id') INT(11)"`
	Receivetime   time.Time `json:"receiveTime" xorm:"comment('领取时间') DATETIME"`
	Receiveprice  int       `json:"receivePrice" xorm:"comment('领取金额，单位：分') INT(11)"`
	Classid       int64     `json:"classId" xorm:"comment('课程id') BIGINT(11)"`
	Activityid    int64     `json:"activityId" xorm:"comment('红包活动id') BIGINT(20)"`
	Status        int       `json:"status" xorm:"default 1 comment('状态 1：正常，0：失效, 2-已提现') TINYINT(3)"`
	Updatetime    time.Time `json:"updateTime" xorm:"comment('更新时间') TIMESTAMP"`
	Putforwardid  int       `json:"putForwardId" xorm:"comment('提现记录ID') INT(11)"`
}
