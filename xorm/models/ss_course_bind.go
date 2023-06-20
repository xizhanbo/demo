package models

import (
	"time"
)

type SsCourseBind struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SsId       int       `json:"ss_id" xorm:"not null comment('双师课程id') index INT(11)"`
	ZxId       int       `json:"zx_id" xorm:"not null comment('在线课程id') index INT(11)"`
	GoodsId    string    `json:"goods_id" xorm:"default '' comment('关联商品id') index VARCHAR(100)"`
	Title      string    `json:"title" xorm:"default '' comment('课程标题') VARCHAR(255)"`
	PayLimit   string    `json:"pay_limit" xorm:"default '' comment('售卖渠道') VARCHAR(100)"`
	Status     int       `json:"status" xorm:"not null default 0 comment('状态 0未上线 1已上线 2已下线') index TINYINT(1)"`
	Updatetime time.Time `json:"updateTime" xorm:"comment('修改时间') DATETIME"`
	Operator   string    `json:"operator" xorm:"default '' comment('操作人') VARCHAR(100)"`
	Createtime time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
}
