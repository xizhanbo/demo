package models

import (
	"time"
)

type MemberClass struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Rid         int       `json:"rid" xorm:"not null comment('课程id') INT(11)"`
	CateId      int       `json:"cate_Id" xorm:"not null comment('java分类id') INT(11)"`
	CategoryId  int       `json:"category_id" xorm:"comment('在线分类id') INT(11)"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('是否删除 0 否 1是') TINYINT(4)"`
	MemberPrice string    `json:"member_price" xorm:"comment('会员价') DECIMAL(10)"`
	CreatedAt   time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	UpdateAt    time.Time `json:"update_at" xorm:"comment('修改时间') DATETIME"`
	Creator     string    `json:"creator" xorm:"comment('操作人') VARCHAR(50)"`
}
