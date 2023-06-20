package models

type DzProvince struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Provinceid string `json:"provinceID" xorm:"VARCHAR(6)"`
	Province   string `json:"province" xorm:"VARCHAR(40)"`
}
