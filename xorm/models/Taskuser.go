package models

type Taskuser struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	NiceName string `json:"nice_name" xorm:"comment('昵称') CHAR(20)"`
	UserName string `json:"user_name" xorm:"comment('蓝色后台用户名') VARCHAR(50)"`
	Issuper  int    `json:"issuper" xorm:"default 0 comment('{0:普通管理员,1超级管理员}') TINYINT(1)"`
	Isdel    int    `json:"isdel" xorm:"default 0 comment('{0:未删除，1:已删除}') TINYINT(1)"`
	UserRole int    `json:"user_role" xorm:"default 1 comment('{1:前端技术，2:后台技术}') TINYINT(1)"`
	Email    string `json:"email" xorm:"comment('邮箱') VARCHAR(255)"`
}
