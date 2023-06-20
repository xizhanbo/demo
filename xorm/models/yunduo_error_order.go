package models

import (
	"time"
)

type YunduoErrorOrder struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	OrderNumber   string    `json:"order_number" xorm:"not null comment('订单编号') unique VARCHAR(100)"`
	UserName      string    `json:"user_name" xorm:"comment('用户名') index VARCHAR(50)"`
	RegisterPhone string    `json:"register_phone" xorm:"comment('注册手机号') index VARCHAR(11)"`
	Phone         string    `json:"phone" xorm:"comment('收获手机号') index VARCHAR(11)"`
	Status        int       `json:"status" xorm:"default 0 comment('状态 0推送失败 1推送成功') TINYINT(1)"`
	ErrorMsg      string    `json:"error_msg" xorm:"comment('失败原因') VARCHAR(255)"`
	CreateTime    time.Time `json:"create_time" xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime    time.Time `json:"update_time" xorm:"default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
	Operator      string    `json:"operator" xorm:"comment('操作人') VARCHAR(50)"`
	PushPhone     string    `json:"push_phone" xorm:"comment('推送客户手机号') VARCHAR(11)"`
	PushId        string    `json:"push_id" xorm:"comment('推送客户ID') VARCHAR(50)"`
	PushManager   string    `json:"push_manager" xorm:"comment('推送客户经理') VARCHAR(50)"`
	AgainTime     time.Time `json:"again_time" xorm:"comment('重推时间') DATETIME"`
	Type          int       `json:"type" xorm:"default 1 comment('工单类型：0退费单，1报名单') TINYINT(1)"`
}
