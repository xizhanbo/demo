package models

import (
	"time"
)

type HignendResultCopy struct {
	Id              int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Username        string    `json:"username" xorm:"default '' comment('用户名') index VARCHAR(50)"`
	AreaName        string    `json:"area_name" xorm:"default '' comment('地区名称') index VARCHAR(50)"`
	CompanyNum      string    `json:"company_num" xorm:"default '' comment('单位id') index VARCHAR(50)"`
	CompanyName     string    `json:"company_name" xorm:"default '' comment('单位名称') index VARCHAR(50)"`
	Name            string    `json:"name" xorm:"default '' comment('学员姓名') index VARCHAR(50)"`
	Admissionticket string    `json:"admissionticket" xorm:"default '' comment('准考证号') index VARCHAR(50)"`
	Phone           string    `json:"phone" xorm:"default '' comment('手机号') CHAR(11)"`
	Sid             string    `json:"sid" xorm:"default '' comment('身份证号') CHAR(18)"`
	Status          int       `json:"status" xorm:"comment('状态 0未支付 1支付') index TINYINT(1)"`
	IsShow          int       `json:"is_show" xorm:"default 0 comment('是否查看（供审核使用） 0未查看 1已查看') TINYINT(1)"`
	BeginAt         time.Time `json:"begin_at" xorm:"comment('入围开始时间') DATETIME"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"comment('更新时间') DATETIME"`
	Operator        string    `json:"Operator" xorm:"default '' comment('审核操作人') VARCHAR(50)"`
}
