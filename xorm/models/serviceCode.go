package models

import (
	"time"
)

type Servicecode struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ChannelType int       `json:"channel_type" xorm:"not null comment('渠道类型：1公众号,2小程序,3专题页,4课程详情,5一键加群,6短信链接') TINYINT(1)"`
	Description string    `json:"description" xorm:"not null comment('说明') TEXT"`
	ServiceCode string    `json:"service_code" xorm:"not null comment('客服码') TEXT"`
	ReceiveCode string    `json:"receive_code" xorm:"not null comment('在线二维码') VARCHAR(100)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('状态:0未上线,1已上线,2已下线,3删除') TINYINT(1)"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null comment('创建时间') DATETIME"`
	CreatorName string    `json:"creator_name" xorm:"not null comment('操作人') VARCHAR(100)"`
	VisitNum    int       `json:"visit_num" xorm:"not null default 0 comment('访问次数') INT(10)"`
	UpdateName  string    `json:"update_name" xorm:"default '' comment('修改人') VARCHAR(100)"`
	UpdateAt    time.Time `json:"update_at" xorm:"comment('修改时间') DATETIME"`
}
