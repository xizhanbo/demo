package models

import (
	"time"
)

type CateTypeLog struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Operater    string    `json:"operater" xorm:"not null comment('操作人') VARCHAR(50)"`
	OperateTime time.Time `json:"operate_time" xorm:"not null comment('操作时间') DATETIME"`
	Content     string    `json:"content" xorm:"not null comment('操作内容') VARCHAR(150)"`
	Type        string    `json:"type" xorm:"not null comment('操作类型') VARCHAR(50)"`
}
