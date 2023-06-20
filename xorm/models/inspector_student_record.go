package models

import (
	"time"
)

type InspectorStudentRecord struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	InspectorId   int       `json:"inspector_id" xorm:"not null default 0 comment('inspector_student_class表id') index INT(11)"`
	VisitType     int       `json:"visit_type" xorm:"not null default 0 comment('回访形式；0其他、1电话、2微信、3微信语音') TINYINT(1)"`
	VisitUsername string    `json:"visit_username" xorm:"not null default '' comment('回访人【Cl_Admin表UserName】') index CHAR(100)"`
	VisitTime     time.Time `json:"visit_time" xorm:"not null default '0000-00-00 00:00:00' comment('回访时间 空为’0000-00-00 00:00:00‘') DATETIME"`
	Serial        int       `json:"serial" xorm:"not null default 1 comment('序号') INT(11)"`
	VisitRecord   string    `json:"visit_record" xorm:"comment('回访记录【不超过500个字】') TEXT"`
	CreateAt      time.Time `json:"create_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateAt      time.Time `json:"update_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}
