package models

import (
	"time"
)

type PerformanceReview struct {
	Id             int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ClassId        int64     `json:"class_id" xorm:"not null comment('课程ID') index BIGINT(20)"`
	Price          string    `json:"price" xorm:"comment('业绩结算金额') DECIMAL(10,2)"`
	Status         int       `json:"status" xorm:"default 0 comment('状态：0待审核1审核通过2审核驳回') TINYINT(1)"`
	IsDel          int       `json:"is_del" xorm:"default 0 comment('是否作废：0否1是') TINYINT(1)"`
	CreateName     string    `json:"create_name" xorm:"comment('创建人') VARCHAR(50)"`
	Reviewer       string    `json:"reviewer" xorm:"comment('审核人') VARCHAR(50)"`
	OrderAsyncTime time.Time `json:"order_async_time" xorm:"comment('订单同步时间') DATETIME"`
	CreateTime     time.Time `json:"create_time" xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateName     string    `json:"update_name" xorm:"comment('更新人') VARCHAR(50)"`
	UpdateTime     time.Time `json:"update_time" xorm:"default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}
