package models

type Learnprocess struct {
	Id       int `json:"id" xorm:"not null pk autoincr INT(8)"`
	Userid   int `json:"userid" xorm:"INT(8)"`
	Addscore int `json:"addScore" xorm:"comment('加分') TINYINT(2)"`
	Subscore int `json:"subScore" xorm:"comment('减分') TINYINT(2)"`
	Addtime  int `json:"addTime" xorm:"comment('创建时间') INT(11)"`
}
