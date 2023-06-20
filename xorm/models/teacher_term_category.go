package models

import (
	"time"
)

type TeacherTermCategory struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	ParentId     int       `json:"parent_id" xorm:"not null default 0 comment('父级ID') INT(11)"`
	ExamType     int       `json:"exam_type" xorm:"not null default 1 comment('考试分类 1资格证，2 招聘') TINYINT(1)"`
	TermCategory int       `json:"term_category" xorm:"not null default 1 comment('分类 1学段、2学科') TINYINT(1)"`
	Title        string    `json:"title" xorm:"not null default '' comment('学段、学科名称') CHAR(6)"`
	Sort         int       `json:"sort" xorm:"not null default 9999 comment('序号越小排序越靠前') INT(11)"`
	UpdateAt     time.Time `json:"update_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	CreateAt     time.Time `json:"create_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}
