package models

import (
	"time"
)

type QaReplyVote struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Qid        int       `json:"qid" xorm:"not null INT(10)"`
	Rid        int       `json:"rid" xorm:"not null INT(10)"`
	Userid     int       `json:"userid" xorm:"not null INT(10)"`
	Createtime time.Time `json:"createtime" xorm:"not null DATETIME"`
}
