package models

type Teacherstudent0 struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Teacherid int    `json:"teacherId" xorm:"not null index INT(11)"`
	Username  string `json:"userName" xorm:"not null default '' index VARCHAR(50)"`
}
