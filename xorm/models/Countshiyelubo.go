package models





type Countshiyelubo struct {

	Id	int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Month	int `json:"month" xorm:"not null index INT(2)"`
	Day	int `json:"day" xorm:"not null INT(2)"`
	2018yeji	float64 `json:"2018yeji" xorm:"not null default 0 DOUBLE"`
	2018renshu	float64 `json:"2018renshu" xorm:"not null default 0 DOUBLE"`
	2019yeji	float64 `json:"2019yeji" xorm:"not null default 0 DOUBLE"`
	2019renshu	float64 `json:"2019renshu" xorm:"not null default 0 DOUBLE"`
	2020yeji	float64 `json:"2020yeji" xorm:"not null default 0 DOUBLE"`
	2020renshu	float64 `json:"2020renshu" xorm:"not null default 0 DOUBLE"`

}


