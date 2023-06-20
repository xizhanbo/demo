package models

import (
	"time"
)

type MemberDiscountClass struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('ID') INT(11)"`
	IsDelete  int       `json:"is_delete" xorm:"not null default 0 comment('是否删除 0否  1是') TINYINT(1)"`
	Rid       int       `json:"rid" xorm:"not null default 0 comment('课程id/合集id') index INT(11)"`
	CateId    string    `json:"cate_id" xorm:"not null default '' comment('分类id，逗号分割的数字字符串') index VARCHAR(255)"`
	Sort      int       `json:"sort" xorm:"not null default 0 comment('排序') INT(11)"`
	Type      int       `json:"type" xorm:"not null default 0 comment('类型  0课程 1合集') TINYINT(4)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') TIMESTAMP"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}
