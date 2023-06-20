package models

type Noticelog struct {
	Id         int `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	Userid     int `json:"userid" xorm:"not null INT(11)"`
	Createtime int `json:"createTime" xorm:"not null INT(11)"`
	Type       int `json:"type" xorm:"not null comment('1 课程提示状态') TINYINT(2)"`
}
