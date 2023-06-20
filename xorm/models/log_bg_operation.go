package models

import (
	"time"
)

type LogBgOperation struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Operationtype int       `json:"operationType" xorm:"not null TINYINT(3)"`
	Content       string    `json:"content" xorm:"not null VARCHAR(1000)"`
	Url           string    `json:"url" xorm:"not null VARCHAR(1000)"`
	Userid        int       `json:"userid" xorm:"not null INT(11)"`
	Username      string    `json:"username" xorm:"not null VARCHAR(50)"`
	Createtime    time.Time `json:"createtime" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}
