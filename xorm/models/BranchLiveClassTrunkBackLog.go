package models

import (
	"time"
)

type Branchliveclasstrunkbacklog struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Netclassid     int       `json:"NetclassId" xorm:"not null index INT(10)"`
	Remark         string    `json:"Remark" xorm:"not null VARCHAR(255)"`
	Creatdate      time.Time `json:"CreatDate" xorm:"not null DATETIME"`
	Optuser        string    `json:"OptUser" xorm:"not null VARCHAR(50)"`
	Branchusername string    `json:"BranchUserName" xorm:"not null comment('分校用户名') index VARCHAR(50)"`
	Isdatail       int       `json:"IsDatail" xorm:"not null default 0 TINYINT(4)"`
}
