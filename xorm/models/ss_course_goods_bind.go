package models

import (
	"time"
)

type SsCourseGoodsBind struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SsBindId int       `json:"ss_bind_id" xorm:"comment('课程id') index(idx) INT(11)"`
	GoodsId  int       `json:"goods_id" xorm:"comment('分校id') index(idx) INT(11)"`
	Operator string    `json:"operator" xorm:"comment('添加人') VARCHAR(255)"`
	Addtime  time.Time `json:"addtime" xorm:"comment('添加时间') DATETIME"`
}
