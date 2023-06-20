package models

type Planzengclass struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Userid  int    `json:"userid" xorm:"INT(8)"`
	Classid string `json:"classid" xorm:"VARCHAR(10)"`
	Time    int    `json:"time" xorm:"comment('时间') INT(10)"`
}
