package models

import (
	"time"
)

type Consultationrecord struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Mobile           string    `json:"mobile" xorm:"comment('咨询用户手机号') index VARCHAR(11)"`
	Examtypeid       int       `json:"examTypeId" xorm:"default 0 comment('考试类型id') INT(1)"`
	Consultationtime time.Time `json:"consultationTime" xorm:"comment('咨询时间') DATETIME"`
	Addtime          time.Time `json:"addTime" xorm:"comment('添加时间') DATETIME"`
	Departmentid     int       `json:"departmentId" xorm:"comment('事业部id') INT(10)"`
	Examtypename     string    `json:"examTypeName" xorm:"comment('考试类型名称') VARCHAR(20)"`
	Coursetitle      string    `json:"courseTitle" xorm:"comment('课程标题') VARCHAR(50)"`
}
