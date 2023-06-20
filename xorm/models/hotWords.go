package models

import (
	"time"
)

type Hotwords struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Word      string    `json:"word" xorm:"not null comment('词') VARCHAR(100)"`
	Sort      int       `json:"sort" xorm:"not null TINYINT(2)"`
	Creator   string    `json:"creator" xorm:"not null comment('创建人') VARCHAR(50)"`
	Createdat time.Time `json:"createdAt" xorm:"not null default CURRENT_TIMESTAMP comment('新增时间') TIMESTAMP"`
	Status    int       `json:"status" xorm:"default 1 comment('1-显示
2-删除') index TINYINT(1)"`
	Deletedat time.Time `json:"deletedAt" xorm:"comment('删除时间') TIMESTAMP"`
	Cate      int       `json:"cate" xorm:"default 1 comment('1-课程
2-资讯') TINYINT(2)"`
	Clicknum int `json:"clickNum" xorm:"default 0 comment('热词点击数量') INT(11)"`
	Viewnum  int `json:"viewNum" xorm:"default 0 comment('通过热词搜索结果进入详情页次数') INT(11)"`
}
