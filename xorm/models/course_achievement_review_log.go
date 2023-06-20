package models

import (
	"time"
)

type CourseAchievementReviewLog struct {
	Id                int64     `json:"id" xorm:"pk comment('id') BIGINT(20)"`
	ClassId           int64     `json:"class_id" xorm:"not null comment('课程ID') index BIGINT(20)"`
	Price             string    `json:"price" xorm:"comment('生效中业绩金额') DECIMAL(10,2)"`
	OldPrice          string    `json:"old_price" xorm:"comment('修改的业绩金额') DECIMAL(10,2)"`
	OrderAsyncTime    time.Time `json:"order_async_time" xorm:"comment('结算金额生效日期') DATETIME"`
	OldOrderAsyncTime time.Time `json:"old_order_async_time" xorm:"comment('结算金额修改日期

') DATETIME"`
	Creator string `json:"creator" xorm:"comment('结算金额修改人

') VARCHAR(50)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('添加日期') DATETIME"`
	Reviewer   string    `json:"reviewer" xorm:"comment('审核人') VARCHAR(50)"`
	AuditTime  time.Time `json:"audit_time" xorm:"comment('审核日期

') DATETIME"`
}
