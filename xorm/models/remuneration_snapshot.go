package models

import (
	"time"
)

type RemunerationSnapshot struct {
	Id         int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(11)"`
	TeacherId  int       `json:"teacher_id" xorm:"not null comment('FK:Teachers表TeacherId,教师/助教ID') index(un_teacher_id_create_time) INT(11)"`
	SujectId   string    `json:"suject_id" xorm:"not null default '' comment('讲授科目信息,以冒号分割进行存储') VARCHAR(15)"`
	Salary     string    `json:"salary" xorm:"not null default 0.00 comment('教师/助教 时薪, 兼职助教 元/课次，其余 元/小时') DECIMAL(10,2)"`
	Operator   string    `json:"operator" xorm:"not null default '' comment('操作人用户名') index VARCHAR(50)"`
	CreateTime time.Time `json:"create_time" xorm:"default CURRENT_TIMESTAMP comment('创建时间') index(un_teacher_id_create_time) DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
	Type       int       `json:"type" xorm:"not null default 1 comment('讲师类型 1讲师 2助教') TINYINT(1)"`
}
