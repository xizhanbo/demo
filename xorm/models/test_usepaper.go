package models

import (
	"time"
)

type TestUsepaper struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Paperid    int       `json:"paperid" xorm:"not null default 0 INT(10)"`
	Papername  string    `json:"paperName" xorm:"not null VARCHAR(255)"`
	Year       int       `json:"year" xorm:"not null SMALLINT(6)"`
	Areaid     string    `json:"areaid" xorm:"VARCHAR(255)"`
	Type       int       `json:"type" xorm:"not null TINYINT(1)"`
	Qids       string    `json:"qids" xorm:"not null VARCHAR(4000)"`
	Createdate time.Time `json:"createdate" xorm:"DATETIME"`
	Status     int       `json:"status" xorm:"not null default 1 TINYINT(1)"`
}
