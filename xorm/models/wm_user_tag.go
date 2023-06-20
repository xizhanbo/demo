package models

import (
	"time"
)

type WmUserTag struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	SrmUserMobile string    `json:"srm_user_mobile" xorm:"not null comment('咨询用户手机号') index VARCHAR(20)"`
	TagId         int       `json:"tag_id" xorm:"not null comment('标签表、分群任务表id') index INT(11)"`
	TagType       int       `json:"tag_type" xorm:"not null comment('标签类别1网图标签2分群标签3知小群_系统标签4知小群_微信标签') TINYINT(1)"`
	TagSource     int       `json:"tag_source" xorm:"default 1 comment('1网图2知小群3星图') TINYINT(1)"`
	Status        int       `json:"status" xorm:"default 0 comment('正常0 删除1') TINYINT(1)"`
	CreatedAt     time.Time `json:"created_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"default CURRENT_TIMESTAMP DATETIME"`
}
