package models

import (
	"time"
)

type LeadsOpportunity struct {
	Id                    int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LeadsId               int       `json:"leads_id" xorm:"not null comment('线索id') index INT(11)"`
	OpportunityId         int       `json:"opportunity_id" xorm:"comment('机会id') INT(11)"`
	RecordId              int64     `json:"record_id" xorm:"not null comment('咨询记录id') BIGINT(20)"`
	Mobile                string    `json:"mobile" xorm:"not null default '' comment('手机号码') index VARCHAR(20)"`
	OpportunityReportTime time.Time `json:"opportunity_report_time" xorm:"comment('线索转化为机会上报时间') index DATETIME"`
	CreateReportTime      time.Time `json:"create_report_time" xorm:"comment('创建订单上报时间') index DATETIME"`
	PayReportTime         time.Time `json:"pay_report_time" xorm:"comment('下单上报时间') index DATETIME"`
	LeadsTime             time.Time `json:"leads_time" xorm:"comment('线索上报时间') index DATETIME"`
	CreatedAt             time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt             time.Time `json:"updated_at" xorm:"DATETIME"`
}
