package models

type DzArea struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Areaid string `json:"areaID" xorm:"index VARCHAR(50)"`
	Area   string `json:"area" xorm:"VARCHAR(60)"`
	Father string `json:"father" xorm:"index VARCHAR(6)"`
}
