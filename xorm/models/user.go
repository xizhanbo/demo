package models

type User struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	NiceName  string `json:"nice_name" xorm:"comment('昵称') CHAR(20)"`
	LoginName string `json:"login_name" xorm:"comment('登录名') CHAR(20)"`
	LastTime  int    `json:"last_time" xorm:"default 0 comment('最近登录时间') INT(11)"`
	LoginPwd  string `json:"login_pwd" xorm:"comment('登录密码') VARCHAR(32)"`
	Isdel     int    `json:"isdel" xorm:"default 0 comment('{0正常,1:删除}') INT(11)"`
	Issuper   int    `json:"issuper" xorm:"default 0 comment('{0:普通管理员,1超级管理员}') INT(11)"`
	UserName  string `json:"user_name" xorm:"comment('蓝色后台用户名') VARCHAR(50)"`
}
