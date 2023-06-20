package models

import (
	"time"
)

type Collagegroup struct {
	Id         int64  `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Classid    int    `json:"classId" xorm:"not null comment('课程ID') index INT(11)"`
	Number     int    `json:"number" xorm:"not null comment('拼团人数') INT(11)"`
	Openuser   string `json:"openUser" xorm:"not null comment('团长用户名') index VARCHAR(50)"`
	Openuserid int    `json:"openUserId" xorm:"not null comment('团长ID') INT(11)"`
	Status     int    `json:"status" xorm:"not null comment('拼团状态, 1未开始（团长下单未支付）2-发起, 3-成功, 4-失败(取消)
') index TINYINT(1)"`
	Startat time.Time `json:"startAt" xorm:"comment('拼团发起时间(状态为发起时间)
') DATETIME"`
	Createdat    time.Time `json:"createdAt" xorm:"not null comment('创建时间') DATETIME"`
	Statusat     time.Time `json:"statusAt" xorm:"comment('拼团成功（失败）时间') DATETIME"`
	Autocancelat time.Time `json:"autoCancelAt" xorm:"comment('失效时间(发起拼团后计算)') index DATETIME"`
	Price        int       `json:"price" xorm:"not null comment('拼团价(分)') index INT(11)"`
	Activityid   int64     `json:"activityId" xorm:"not null comment('活动ID') index BIGINT(20)"`
	Sendnotice   int       `json:"sendNotice" xorm:"default 1 comment('距结束六小时是否发送通知：1 发送 0 不发送') index TINYINT(1)"`
}
