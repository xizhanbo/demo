package models

import (
	"time"
)

type XetClass struct {
	Rid              int       `json:"rid" xorm:"not null pk autoincr comment('表ID') INT(11)"`
	NetClassesRid    string    `json:"net_classes_rid" xorm:"not null default '' comment('本地课程rid，NetClasses表rid，逗号分隔的字符串，逗号结尾') VARCHAR(255)"`
	Id               string    `json:"id" xorm:"not null default '' comment('小鹅通ID 字符串格式') unique VARCHAR(255)"`
	ResourceType     int       `json:"resource_type" xorm:"not null default 0 comment('商品类型 1.图文 2.音频 3.视频 4.直播 6.专栏 20.电子书') TINYINT(3)"`
	Title            string    `json:"title" xorm:"not null default '' comment('title') VARCHAR(255)"`
	ImgUrl           string    `json:"img_url" xorm:"not null default '' comment('img_url') VARCHAR(255)"`
	ImgUrlCompressed string    `json:"img_url_compressed" xorm:"not null default '' comment('封面压缩图片url地址') VARCHAR(255)"`
	Price            int       `json:"price" xorm:"not null default 0 comment('价格（分）') INT(11)"`
	State            int       `json:"state" xorm:"not null default 0 comment('状态： 0.有效 1.隐藏') TINYINT(1)"`
	IsStopSell       int       `json:"is_stop_sell" xorm:"not null default 0 comment('停售状态 0:未停售 1：已停售') TINYINT(1)"`
	PublishTime      time.Time `json:"publish_time" xorm:"not null default '1970-01-01 00:00:00' comment('小鹅通的课程上架时间') DATETIME"`
	CreatedTime      time.Time `json:"created_time" xorm:"not null default '1970-01-01 00:00:00' comment('小鹅通的课程创建时间') DATETIME"`
	CreatedAt        time.Time `json:"created_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}
