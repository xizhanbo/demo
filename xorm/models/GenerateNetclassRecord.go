package models

type Generatenetclassrecord struct {
	Id         int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Categoryid int `json:"CategoryId" xorm:"not null comment('考试类型') INT(11)"`
	Provinceid int `json:"ProvinceId" xorm:"not null comment('省份') INT(11)"`
	Coursetype int `json:"Coursetype" xorm:"not null comment('笔面试') INT(11)"`
	Netclassid int `json:"NetclassId" xorm:"not null comment('课程id') INT(11)"`
}
