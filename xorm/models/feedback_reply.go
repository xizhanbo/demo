package models

import (
	"time"
)

type FeedbackReply struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Replycontent  string    `json:"ReplyContent" xorm:"not null comment('回复内容') TEXT"`
	Replyuserid   int       `json:"ReplyUserId" xorm:"not null comment('回复人ID') INT(11)"`
	Replyusername string    `json:"ReplyUserName" xorm:"not null comment('回复人用户名') VARCHAR(50)"`
	Feedbackid    int       `json:"FeedBackId" xorm:"not null INT(11)"`
	Createdate    time.Time `json:"CreateDate" xorm:"not null DATETIME"`
	Status        int       `json:"Status" xorm:"default 0 comment('回复状态') TINYINT(4)"`
}
