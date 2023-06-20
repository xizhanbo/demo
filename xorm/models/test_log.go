package models

import (
	"time"
)

type TestLog struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Userid           int       `json:"userid" xorm:"comment('用户ID') INT(10)"`
	Username         string    `json:"username" xorm:"comment('用户名') VARCHAR(50)"`
	Lessionid        int       `json:"lessionid" xorm:"comment('课件(试卷)ID') INT(10)"`
	Answer           string    `json:"answer" xorm:"comment('学员答案') VARCHAR(1000)"`
	Rightanswercount int       `json:"rightAnswerCount" xorm:"comment('试题答对个数') INT(5)"`
	Usetime          string    `json:"usetime" xorm:"comment('答题用时') VARCHAR(20)"`
	Createtime       time.Time `json:"createtime" xorm:"comment('提交答案时间') DATETIME"`
	Allanswercount   int       `json:"allAnswerCount" xorm:"comment('该套试卷题目个数') INT(5)"`
}
