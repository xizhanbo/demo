package models

import (
	"time"
)

type ClUserInfo struct {
	Id              int       `json:"ID" xorm:"not null pk autoincr comment('自增id') INT(10)"`
	Username        string    `json:"UserName" xorm:"not null comment('用户名') unique VARCHAR(50)"`
	Province        string    `json:"Province" xorm:"comment('省份') VARCHAR(50)"`
	City            string    `json:"City" xorm:"comment('城市') VARCHAR(50)"`
	Countyarea      string    `json:"CountyArea" xorm:"comment('县区') VARCHAR(50)"`
	Detailedaddress string    `json:"DetailedAddress" xorm:"comment('详细地址') VARCHAR(50)"`
	Equipment       string    `json:"Equipment" xorm:"comment('操作设备') VARCHAR(20)"`
	Systemversion   string    `json:"SystemVersion" xorm:"comment('系统版本号') VARCHAR(20)"`
	Appversion      string    `json:"AppVersion" xorm:"comment('APP版本号') VARCHAR(20)"`
	Joindate        time.Time `json:"JoinDate" xorm:"not null comment('注册时间') DATETIME"`
	Fromuser        int       `json:"FromUser" xorm:"comment('来源') INT(10)"`
	Updatetime      time.Time `json:"UpdateTime" xorm:"comment('更新时间') DATETIME"`
}
