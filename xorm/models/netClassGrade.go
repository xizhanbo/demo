package models

import (
	"time"
)

type Netclassgrade struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	GradeName          string    `json:"grade_name" xorm:"not null comment('班级名称') VARCHAR(120)"`
	NetClassCategoryId int64     `json:"net_class_category_id" xorm:"not null comment('考试id') index BIGINT(20)"`
	FirstSubject       int64     `json:"first_subject" xorm:"comment('一级科目id') BIGINT(20)"`
	SubjectType        int64     `json:"subject_type" xorm:"comment('二级科目') BIGINT(20)"`
	SubjectStatus      int       `json:"subject_status" xorm:"not null comment('1笔试 2面试 3纯图书 4综合') INT(1)"`
	ShowStatus         int       `json:"show_status" xorm:"not null comment('1显示 2隐藏') INT(1)"`
	CreateTime         time.Time `json:"create_time" xorm:"not null comment('创建时间') index DATETIME"`
	OperaterId         int64     `json:"operater_id" xorm:"not null comment('创建人') index BIGINT(20)"`
	OperaterName       string    `json:"operater_name" xorm:"not null comment('创建人昵称') index VARCHAR(255)"`
	Status             int       `json:"status" xorm:"not null default 1 comment('1 普通 4删除') INT(1)"`
	Intro              string    `json:"intro" xorm:"comment('班别描述') VARCHAR(255)"`
	Videotype          int       `json:"videoType" xorm:"default 0 comment('班次类型') INT(11)"`
}
