package models

type Headuserclass struct {
	Id          int    `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Username    string `json:"userName" xorm:"comment('用户名称') VARCHAR(50)"`
	Qqnum       string `json:"QQNum" xorm:"comment('QQ号') VARCHAR(10)"`
	Rid         int    `json:"rid" xorm:"comment('课程ID') INT(10)"`
	Headuser    string `json:"headUser" xorm:"comment('班主任名称') VARCHAR(50)"`
	Desc        string `json:"desc" xorm:"comment('备注') VARCHAR(200)"`
	IsAddGroups int    `json:"is_add_groups" xorm:"default 0 comment('是否加群 1:加群，0:未加群 ，') TINYINT(1)"`
}
