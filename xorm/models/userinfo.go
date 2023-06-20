package models

import (
	"time"
)

type Userinfo struct {
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键，自增') INT(10)"`
	Mobile         string    `json:"mobile" xorm:"comment('唯一，手机号') unique VARCHAR(20)"`
	Loginip        string    `json:"loginIp" xorm:"comment('登录IP') VARCHAR(50)"`
	Username       string    `json:"username" xorm:"comment('用户名') VARCHAR(50)"`
	Mobileprovince string    `json:"mobileProvince" xorm:"comment('手机号归属地省份') VARCHAR(50)"`
	Mobilecity     string    `json:"mobileCity" xorm:"comment('手机号归属地市') VARCHAR(50)"`
	Ipprovince     string    `json:"ipProvince" xorm:"comment('IP所在省') VARCHAR(50)"`
	Ipcity         string    `json:"ipCity" xorm:"comment('IP所在市') VARCHAR(50)"`
	Carrier        string    `json:"carrier" xorm:"comment('运营商') VARCHAR(50)"`
	Lastlogintime  time.Time `json:"lastLoginTime" xorm:"comment('上次登录时间') DATETIME"`
}
