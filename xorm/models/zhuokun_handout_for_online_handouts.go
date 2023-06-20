package models

import (
	"time"
)

type ZhuokunHandoutForOnlineHandouts struct {
	Id               int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	ZhuokunHandoutId int64     `json:"zhuokun_handout_id" xorm:"not null comment('卓坤医疗资料ID') index BIGINT(20)"`
	HandoutId        int64     `json:"handout_id" xorm:"not null comment('华图在线讲义ID') index BIGINT(20)"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"DATETIME"`
	CreatedAt        time.Time `json:"created_at" xorm:"DATETIME"`
}
