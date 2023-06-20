package models

import (
	"time"
)

type WinningRecord struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Userid    int       `json:"UserId" xorm:"not null comment('用户ID') INT(11)"`
	Username  string    `json:"UserName" xorm:"comment('用户名称') VARCHAR(50)"`
	Chaitime  time.Time `json:"ChaiTime" xorm:"comment('获奖时间') DATETIME"`
	Prizetype string    `json:"prizeType" xorm:"comment('奖品名称') VARCHAR(50)"`
	Pici      string    `json:"pici" xorm:"comment('活动名称') VARCHAR(10)"`
	Riqi      string    `json:"riqi" xorm:"comment('获奖日期') VARCHAR(20)"`
}
