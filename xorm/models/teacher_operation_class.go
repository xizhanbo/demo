package models

import (
	"time"
)

type TeacherOperationClass struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TagType   int       `json:"tag_type" xorm:"not null comment('标签类型：1资格笔试，2资格面试，3招聘笔试，4招聘面试') index(_tag) TINYINT(4)"`
	TagId     int       `json:"tag_id" xorm:"not null comment('标签id') index(_tag) INT(10)"`
	Province  int       `json:"province" xorm:"not null default 0 comment('省份：0全部、1000全国') index(_tag) INT(10)"`
	Term      int       `json:"term" xorm:"not null default 0 comment('学段：0全部') index(_tag) TINYINT(3)"`
	Category  int       `json:"category" xorm:"not null default 0 comment('学科：0全部') index(_tag) INT(11)"`
	ItemType  int       `json:"item_type" xorm:"not null comment('类型，0课程，1合集，2集合') TINYINT(3)"`
	ItemId    int64     `json:"item_id" xorm:"not null comment('课程、合集、集合id') BIGINT(20)"`
	ExamId    int       `json:"exam_id" xorm:"not null default 0 comment('课程、合集、集合的考试类型ID') INT(10)"`
	ItemOrder int       `json:"item_order" xorm:"not null comment('顺序，从1开始') INT(10)"`
	Operator  string    `json:"operator" xorm:"not null comment('操作人') VARCHAR(128)"`
	Status    int       `json:"status" xorm:"not null default 1 comment('状态：0未上线或课程关闭，1已上线，2已下线，3已删除') TINYINT(4)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}
