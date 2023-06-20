package models

type Newmooddiary struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(8)"`
	Userid     int    `json:"userid" xorm:"default 0 INT(8)"`
	Username   string `json:"username" xorm:"default '0' VARCHAR(255)"`
	Content    string `json:"content" xorm:"VARCHAR(255)"`
	Categoryid int    `json:"categoryId" xorm:"TINYINT(2)"`
	Provinceid int    `json:"provinceId" xorm:"TINYINT(2)"`
	Addtime    int    `json:"addTime" xorm:"default 0 INT(10)"`
	Status     int    `json:"status" xorm:"default 1 comment('1 审核通过') TINYINT(1)"`
}
