package models

import (
	"time"
)

type AuditRecord struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ProxyId     int64     `json:"proxy_id" xorm:"not null comment('代理id') index BIGINT(20)"`
	AuditorName string    `json:"auditor_name" xorm:"not null comment('审核人用户名') VARCHAR(20)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	Status      int       `json:"status" xorm:"not null comment('审核状态') index TINYINT(1)"`
	ProxyName   string    `json:"proxy_name" xorm:"not null comment('合伙人名字') VARCHAR(20)"`
}
