package models

type Lessioninfo struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Rid     int    `json:"rid" xorm:"index INT(11)"`
	Fileurl string `json:"fileUrl" xorm:"VARCHAR(350)"`
}
