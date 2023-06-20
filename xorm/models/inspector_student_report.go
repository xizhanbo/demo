package models

import (
	"time"
)

type InspectorStudentReport struct {
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	InspectorId     int       `json:"inspector_id" xorm:"not null default 0 comment('inspector_student_class表的id') index INT(11)"`
	Serial          int       `json:"serial" xorm:"not null default 1 comment('序号') INT(11)"`
	ReportTime      time.Time `json:"report_time" xorm:"not null comment('生成报告时间【发起生成报告提交的时间，并非报告的实际生成时间】') DATETIME"`
	ReportStartDate time.Time `json:"report_start_date" xorm:"not null comment('生成报告的时间段-开始') DATE"`
	ReportEndDate   time.Time `json:"report_end_date" xorm:"not null comment('生成报告的时间段-结束') DATE"`
	Operator        string    `json:"operator" xorm:"not null default '' comment('生成人【创建工单的后台用户真实姓名】') index CHAR(100)"`
	ReportUrl       string    `json:"report_url" xorm:"not null default '' comment('报告地址') VARCHAR(255)"`
	Review          string    `json:"review" xorm:"comment('督学点评【不超过500字】') TEXT"`
	UpdateAt        time.Time `json:"update_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	CreateAt        time.Time `json:"create_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}
