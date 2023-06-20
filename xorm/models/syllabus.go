package models

import (
	"time"
)

type Syllabus struct {
	Id            int       `json:"id" xorm:"not null pk autoincr index(test_index) INT(10)"`
	ParentId      int       `json:"parent_id" xorm:"default 0 comment('父id') index INT(11)"`
	ClassId       int       `json:"class_id" xorm:"default 0 comment('课程表id') index index(union_index) INT(11)"`
	CoursewareId  int       `json:"courseware_id" xorm:"default 0 comment('课件id') index INT(11)"`
	Name          string    `json:"name" xorm:"not null comment('阶段名称/课程名称/课件名称') index VARCHAR(200)"`
	SerialNumber  int       `json:"serial_number" xorm:"default 0 comment('排序，越小越靠前') index INT(10)"`
	ClassHour     int       `json:"class_hour" xorm:"default 0 comment('课时(分钟)') index INT(11)"`
	Type          int       `json:"type" xorm:"not null default 0 comment('0阶段1课程2课件') index(test_index) index TINYINT(1)"`
	NetClassId    int       `json:"net_class_id" xorm:"not null comment('大纲绑定的课程id') index index(path) INT(11)"`
	IsTrial       int       `json:"is_trial" xorm:"default 0 comment('是否是试听课件0否1是') TINYINT(1)"`
	VideoType     int       `json:"video_type" xorm:"default 0 comment('课件类型1点播2直播3直播回放4阶段测试题5音频') index(union_index) index TINYINT(1)"`
	IsPass        int       `json:"is_pass" xorm:"default 1 comment('是否审核通过1已通过2已驳回') index(test_index) index(union_index) TINYINT(1)"`
	IsDel         int       `json:"is_del" xorm:"default 0 comment('是否已删除0否1是') index(test_index) index(union_index) TINYINT(1)"`
	Hours         int       `json:"hours" xorm:"default 0 comment('课时') index INT(11)"`
	CoursewareNum int       `json:"courseware_num" xorm:"default 0 comment('课件顺序号') index INT(11)"`
	IsMidified    int       `json:"is_midified" xorm:"default 0 comment('是否修改过节点名称0否1是') TINYINT(1)"`
	CreatedAt     time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"DATETIME"`
	Path          string    `json:"path" xorm:"index(path) VARCHAR(255)"`
}
