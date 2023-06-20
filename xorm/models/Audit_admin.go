package models

type AuditAdmin struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserName string `json:"user_name" xorm:"comment('管理员用户名') VARCHAR(20)"`
	Password string `json:"password" xorm:"comment('管理员密码') VARCHAR(32)"`
}
