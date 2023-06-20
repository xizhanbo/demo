package models

import (
	"time"
)

type XxkApplyLog struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ApplyOrder  string    `json:"apply_order" xorm:"not null default '0' comment('申请单编号') CHAR(20)"`
	ApplyStatus int       `json:"apply_status" xorm:"not null default 0 comment('申请单状态') TINYINT(4)"`
	Operator    string    `json:"operator" xorm:"not null default '' comment('执行操作的用户') VARCHAR(20)"`
	Comment     string    `json:"comment" xorm:"not null default '' comment('意见') VARCHAR(50)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('操作时间') TIMESTAMP"`
}
