package models

import (
	"time"
)

type YunduoBillRules struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	DepartmentId   string    `json:"department_id" xorm:"not null default '' comment('云朵二级部门id') index VARCHAR(32)"`
	DepartmentName string    `json:"department_name" xorm:"not null default '' comment('云朵二级部门名称') VARCHAR(64)"`
	Rule           string    `json:"rule" xorm:"not null comment('规则') TEXT"`
	Status         int       `json:"status" xorm:"not null default 0 comment('上下线状态，0禁用，1启用') TINYINT(4)"`
	Operator       string    `json:"operator" xorm:"not null default '' comment('最新维护人') VARCHAR(64)"`
	OperateAt      time.Time `json:"operate_at" xorm:"not null comment('最新维护时间') DATETIME"`
}
