package models

import (
	"time"
)

type SsCourse struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Courseid             int       `json:"courseId" xorm:"not null comment('神一课程id') INT(11)"`
	Code                 string    `json:"code" xorm:"not null default '' comment('课程编号') VARCHAR(100)"`
	Name                 string    `json:"name" xorm:"not null default '' comment('课程标题') VARCHAR(255)"`
	Productrespdtocode   string    `json:"productRespDtoCode" xorm:"default '' comment('所属考试编码') VARCHAR(100)"`
	Productrespdtoname   string    `json:"productRespDtoName" xorm:"default '' comment('所属考试类型名称') VARCHAR(255)"`
	Examtypeid           int       `json:"examTypeId" xorm:"comment('考试类型id') INT(11)"`
	Examtypename         string    `json:"examTypeName" xorm:"default '' comment('考试类型名称') VARCHAR(100)"`
	Branchdepartmentid   int       `json:"branchDepartmentId" xorm:"comment('分校id') INT(11)"`
	Branchdepartmentname string    `json:"branchDepartmentName" xorm:"default '' comment('分校名称') VARCHAR(255)"`
	Createtime           string    `json:"createTime" xorm:"default '' comment('创建时间') VARCHAR(13)"`
	Validitydatestart    string    `json:"validityDateStart" xorm:"default '' comment('课程有效期开始') VARCHAR(13)"`
	Validitydateend      string    `json:"validityDateEnd" xorm:"default '' comment('课程有效期结束') VARCHAR(13)"`
	Ispeoplenumberlimit  int       `json:"isPeopleNumberLimit" xorm:"default 0 comment('是否有人数上线 1有0没有') TINYINT(1)"`
	Maxpeoplenumber      int       `json:"maxPeopleNumber" xorm:"default 0 comment('最高人数') INT(11)"`
	Status               int       `json:"status" xorm:"default 1 comment('课程状态 1默认上线 0：下线') TINYINT(1)"`
	Updatetime           time.Time `json:"updateTime" xorm:"comment('修改时间') DATETIME"`
}
