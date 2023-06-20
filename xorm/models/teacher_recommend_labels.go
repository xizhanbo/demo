package models

import (
	"time"
)

type TeacherRecommendLabels struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	TagName      string    `json:"tag_name" xorm:"not null default '' comment('标签名称') VARCHAR(32)"`
	Type         int       `json:"type" xorm:"not null default 0 comment('标签类型 1资格笔试 2资格面试 3招聘笔试 4招聘面试') TINYINT(1)"`
	Attribute    int       `json:"attribute" xorm:"not null default 0 comment('标签属性 1运营 2产品') TINYINT(1)"`
	OperatorName string    `json:"operator_name" xorm:"not null default '' comment('操作人') VARCHAR(128)"`
	Sort         int       `json:"sort" xorm:"not null default 1 comment('排序【规则-正序】') INT(11)"`
	CreatedAt    time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"comment('修改时间') DATETIME"`
	DeletedAt    time.Time `json:"deleted_at" xorm:"comment('删除时间') DATETIME"`
}
