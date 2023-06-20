package models

import (
	"time"
)

type Searchhistory struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Username  string    `json:"userName" xorm:"comment('用户名') index VARCHAR(100)"`
	Word      string    `json:"word" xorm:"comment('搜索词') VARCHAR(200)"`
	Status    int       `json:"status" xorm:"default 1 comment('状态1-正常, 2-失效') index TINYINT(1)"`
	Createdat time.Time `json:"createdAt" xorm:"comment('创建时间') TIMESTAMP"`
	Updateat  time.Time `json:"updateAt" xorm:"default CURRENT_TIMESTAMP comment('最后更新时间') TIMESTAMP"`
	Cate      int       `json:"cate" xorm:"default 2 comment('分类, 1-课程,2-分类') index TINYINT(2)"`
}
