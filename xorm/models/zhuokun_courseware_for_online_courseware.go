package models

import (
	"time"
)

type ZhuokunCoursewareForOnlineCourseware struct {
	Id                  int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ZhuokunCoursewareId int64     `json:"zhuokun_courseware_id" xorm:"not null comment('卓坤课件ID') index BIGINT(20)"`
	CoursewareId        int64     `json:"courseware_id" xorm:"not null comment('华图课件ID') index BIGINT(20)"`
	Type                int       `json:"type" xorm:"not null comment('课件类型：1直播课件、2录播课件') TINYINT(3)"`
	BjyPartnerId        int64     `json:"bjy_partner_id" xorm:"not null comment('卓坤医疗百家云账号ID') BIGINT(20)"`
	BjyPartnerKey       string    `json:"bjy_partner_key" xorm:"not null comment('卓坤医疗百家云账号key') VARCHAR(255)"`
	UpdatedAt           time.Time `json:"updated_at" xorm:"DATETIME"`
	CreatedAt           time.Time `json:"created_at" xorm:"DATETIME"`
}
