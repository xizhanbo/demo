package models

type ClPermit struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Userid   int    `json:"UserId" xorm:"not null INT(11)"`
	Username int    `json:"UserName" xorm:"INT(11)"`
	Permit   string `json:"Permit" xorm:"VARCHAR(800)"`
	Role     string `json:"Role" xorm:"VARCHAR(800)"`
}
