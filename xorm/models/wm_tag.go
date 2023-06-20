package models

import (
	"time"
)

type WmTag struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OutsideTagId int       `json:"outside_tag_id" xorm:"default 0 comment('外部标签id') index INT(11)"`
	Name         string    `json:"name" xorm:"not null comment('名称') index VARCHAR(100)"`
	ParentId     int       `json:"parent_id" xorm:"not null default 0 comment('父节点id') index INT(11)"`
	Editor       string    `json:"editor" xorm:"not null comment('编辑人') VARCHAR(50)"`
	Founder      string    `json:"founder" xorm:"not null comment('创建人') VARCHAR(50)"`
	ExamType     string    `json:"exam_type" xorm:"comment('考试类型') TEXT"`
	UserNum      int64     `json:"user_num" xorm:"not null default 0 comment('标签用户量') BIGINT(20)"`
	TagType      int       `json:"tag_type" xorm:"not null default 1 comment('1网图标签2分群标签3知小群_系统标签4知小群_微信标签5事件标签') TINYINT(1)"`
	Remarks      string    `json:"remarks" xorm:"comment('备注') TEXT"`
	Status       int       `json:"status" xorm:"not null default 0 comment('0正常1删除') TINYINT(1)"`
	CreatedAt    time.Time `json:"created_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
}
