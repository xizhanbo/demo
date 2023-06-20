package models

import (
	"time"
)

type CourseCardAttribute struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	CardId      int       `json:"card_id" xorm:"not null comment('课程卡ID') index INT(11)"`
	ScopeId     int       `json:"scope_id" xorm:"not null comment('适用范围') index INT(11)"`
	IssueNumber int       `json:"issue_number" xorm:"not null default 0 comment('发行数量') INT(11)"`
	CreateTime  time.Time `json:"create_time" xorm:"comment('生成时间') DATETIME"`
	Creator     string    `json:"creator" xorm:"comment('创建人') VARCHAR(50)"`
	StartTime   time.Time `json:"start_time" xorm:"comment('有效期开始时间') DATETIME"`
	EndTime     time.Time `json:"end_time" xorm:"comment('有效期结束时间') DATETIME"`
}
