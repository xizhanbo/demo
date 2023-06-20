package models

import (
	"time"
)

type TeacherRemuneration struct {
	Id                  int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(11)"`
	TeacherId           int       `json:"teacher_id" xorm:"not null comment('FK:Teachers表TeacherId,教师/助教ID') index INT(11)"`
	SujectId            string    `json:"suject_id" xorm:"not null default '' comment('讲授科目信息,以冒号分割进行存储') index VARCHAR(15)"`
	Salary              string    `json:"salary" xorm:"not null default 0.00 comment('教师/助教 时薪, 兼职助教 元/课次，其余 元/小时') DECIMAL(10,2)"`
	MaintenanceTime     time.Time `json:"maintenance_time" xorm:"default '1970-01-01 00:00:00' comment('薪酬维护时间') DATETIME"`
	MaintenanceUser     string    `json:"maintenance_user" xorm:"not null comment('薪酬维护人') index VARCHAR(50)"`
	MaintenanceRealname string    `json:"maintenance_realname" xorm:"not null comment('冗余字段，薪酬维护人真实姓名') VARCHAR(150)"`
	Status              int       `json:"status" xorm:"not null default 0 comment('状态 0有效，1无效') TINYINT(1)"`
	CreateTime          time.Time `json:"create_time" xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime          time.Time `json:"update_time" xorm:"default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
	Type                int       `json:"type" xorm:"not null default 1 comment('讲师类型 1讲师 2助教') TINYINT(1)"`
}
