package models

import (
	"time"
)

type MessageLog struct {
	Tid        int       `json:"tid" xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Optuser    string    `json:"OptUser" xorm:"comment('发起人') VARCHAR(255)"`
	Department string    `json:"department" xorm:"VARCHAR(255)"`
	Topic      string    `json:"topic" xorm:"VARCHAR(255)"`
	Topicid    int       `json:"topicId" xorm:"comment('短信主题id') INT(11)"`
	Messagenum int       `json:"messageNum" xorm:"comment('短信数量') INT(11)"`
	Catalog    string    `json:"catalog" xorm:"comment('目录') VARCHAR(255)"`
	Senddate   time.Time `json:"sendDate" xorm:"comment('发起时间') DATETIME"`
}
