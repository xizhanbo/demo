package models

import (
	"time"
)

type CourseActivationCode struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('激活码主键ID') INT(10)"`
	CardId      int       `json:"card_id" xorm:"not null comment('课程卡ID') index INT(11)"`
	AttributeId int       `json:"attribute_id" xorm:"not null comment('课程卡属性ID') INT(11)"`
	ScopeId     int       `json:"scope_id" xorm:"not null comment('适用范围') INT(11)"`
	CardNum     string    `json:"card_num" xorm:"not null comment('课程卡卡号') unique index VARCHAR(50)"`
	CardPass    string    `json:"card_pass" xorm:"not null comment('课程卡密码') index VARCHAR(50)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('激活状态（0:未激活 1:已激活 2:已失效）') index TINYINT(3)"`
	CreateTime  time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UserId      int64     `json:"user_id" xorm:"comment('激活用户ID') BIGINT(20)"`
	UserName    string    `json:"user_name" xorm:"comment('激活用户名') VARCHAR(50)"`
	UpdateTime  time.Time `json:"update_time" xorm:"comment('激活时间') DATETIME"`
	Operation   string    `json:"operation" xorm:"not null comment('操作人') VARCHAR(50)"`
	OrderIds    string    `json:"order_ids" xorm:"comment('课程卡激活课程订单ID，多个逗号隔开') TEXT"`
	Creator     string    `json:"creator" xorm:"comment('创建人') VARCHAR(50)"`
	Terminal    int       `json:"terminal" xorm:"comment('来源') TINYINT(3)"`
	StartTime   time.Time `json:"start_time" xorm:"comment('有效期开始时间') index DATETIME"`
	EndTime     time.Time `json:"end_time" xorm:"comment('有效期结束时间') index DATETIME"`
}
