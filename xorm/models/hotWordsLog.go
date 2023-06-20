package models

import (
	"time"
)

type Hotwordslog struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Operator string    `json:"Operator" xorm:"not null comment('操作人') VARCHAR(50)"`
	Action   string    `json:"action" xorm:"not null comment('操作动作') VARCHAR(20)"`
	Hotwords string    `json:"hotWords" xorm:"not null comment('操作热词') VARCHAR(100)"`
	Cate     int       `json:"cate" xorm:"default 1 comment('1-课程\n2-资讯') TINYINT(2)"`
	Createat time.Time `json:"createAt" xorm:"not null default CURRENT_TIMESTAMP comment('添加时间') TIMESTAMP"`
}
