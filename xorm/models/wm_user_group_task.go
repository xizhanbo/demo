package models

import (
	"time"
)

type WmUserGroupTask struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name       string    `json:"name" xorm:"not null comment('分群名称') index VARCHAR(200)"`
	Describe   string    `json:"describe" xorm:"comment('分群描述') TEXT"`
	GroupNum   int64     `json:"group_num" xorm:"not null default 0 comment('用户数') BIGINT(20)"`
	Founder    string    `json:"founder" xorm:"not null comment('创建人') VARCHAR(50)"`
	Editor     string    `json:"editor" xorm:"not null comment('编辑人') VARCHAR(50)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('1未开始2进行中3已完成4已删除') TINYINT(1)"`
	Conditions string    `json:"conditions" xorm:"comment('分群条件') TEXT"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
	Filename   string    `json:"fileName" xorm:"comment('检索结果文件名称') index VARCHAR(100)"`
}
