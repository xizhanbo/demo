package models

import (
	"time"
)

type GoodTeacherClass struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Rid        int       `json:"rid" xorm:"not null default 0 comment('课程id/合集id') index INT(11)"`
	CateId     string    `json:"cate_id" xorm:"not null default '0' comment('分类id') index VARCHAR(255)"`
	Sort       int       `json:"sort" xorm:"not null default 0 comment('排序') INT(11)"`
	Type       int       `json:"type" xorm:"not null comment('类型  0课程 1合集') index TINYINT(4)"`
	Creator    string    `json:"creator" xorm:"not null comment('创建人') VARCHAR(50)"`
	CreatedAt  time.Time `json:"created_at" xorm:"not null comment('创建时间') DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null comment('修改时间') DATETIME"`
	IsDel      int       `json:"is_del" xorm:"not null default 0 comment('是否删除 0否  1是') TINYINT(4)"`
	UpdateName string    `json:"update_name" xorm:"not null comment('修改人') VARCHAR(50)"`
}
