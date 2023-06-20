package models

type DzCity struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Cityid string `json:"cityID" xorm:"index VARCHAR(6)"`
	City   string `json:"city" xorm:"VARCHAR(50)"`
	Father string `json:"father" xorm:"index VARCHAR(6)"`
}
