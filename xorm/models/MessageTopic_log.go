package models

import (
	"time"
)

type MessagetopicLog struct {
	Logid          int       `json:"Logid" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Edittime       time.Time `json:"edittime" xorm:"comment('修改时间') DATETIME"`
	Department     int       `json:"department" xorm:"comment('所属部门') TINYINT(4)"`
	Topic          string    `json:"topic" xorm:"comment('主题名') VARCHAR(255)"`
	Messagetemplet string    `json:"messagetemplet" xorm:"comment('主题内容') VARCHAR(255)"`
	Optuser        string    `json:"OptUser" xorm:"comment('操作员') VARCHAR(255)"`
	Tid            int       `json:"tid" xorm:"comment('主题id') INT(11)"`
}
