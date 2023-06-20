package models

import (
	"time"
)

type TestPaper struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Papername  string    `json:"paperName" xorm:"not null VARCHAR(255)"`
	Year       int       `json:"year" xorm:"not null SMALLINT(6)"`
	Areaid     string    `json:"areaid" xorm:"default '' VARCHAR(255)"`
	Type       int       `json:"type" xorm:"not null TINYINT(1)"`
	Createdate time.Time `json:"createdate" xorm:"DATETIME"`
	Status     int       `json:"status" xorm:"not null default 1 TINYINT(1)"`
}
