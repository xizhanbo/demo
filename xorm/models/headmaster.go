package models

import (
	"time"
)

type Headmaster struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('班主任主键ID') INT(10)"`
	Name               string    `json:"name" xorm:"not null comment('姓名') index VARCHAR(20)"`
	Username           string    `json:"userName" xorm:"not null comment('用户名') VARCHAR(20)"`
	Mobilephone        string    `json:"mobilePhone" xorm:"not null comment('班主任手机号') index VARCHAR(11)"`
	Idcard             string    `json:"idCard" xorm:"comment('身份证号') CHAR(18)"`
	Qqnum              string    `json:"QQNum" xorm:"comment('QQ号') VARCHAR(20)"`
	Creator            string    `json:"creator" xorm:"comment('创建人') VARCHAR(20)"`
	Createdate         time.Time `json:"createDate" xorm:"comment('创建时间') DATETIME"`
	Status             int       `json:"status" xorm:"default 1 comment('1上线，0下线  状态') TINYINT(1)"`
	Smallclassnum      int       `json:"smallClassNum" xorm:"default 0 comment('总维护小总数') index INT(10)"`
	Studentstotal      int       `json:"studentsTotal" xorm:"not null default 0 comment('学生总数') index INT(10)"`
	Maintainqqgroup    int       `json:"maintainQQgroup" xorm:"default 0 comment('维护QQ群总数') index INT(10)"`
	Nowsmallnum        int       `json:"nowSmallNum" xorm:"default 0 comment('当前维护小班数') index INT(10)"`
	Nowstudenstotal    int       `json:"nowStudensTotal" xorm:"default 0 comment('当前学生人数') index INT(10)"`
	Nowmaintainqqgroup int       `json:"nowMaintainQQgroup" xorm:"default 0 comment('当前维护QQ群人数') index INT(10)"`
	Updatedate         time.Time `json:"UpdateDate" xorm:"DATETIME"`
	Rating             string    `json:"rating" xorm:"comment('班主任评级') VARCHAR(10)"`
}
