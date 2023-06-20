package models

import (
	"time"
)

type NetclasscategoryidRecommend struct {
	Id                  int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Rid                 int       `json:"rid" xorm:"not null comment('课程ID') index INT(11)"`
	Categoryid          int       `json:"categoryid" xorm:"not null default 0 comment('推荐课程类型') index TINYINT(4)"`
	Recommenddate       time.Time `json:"recommenddate" xorm:"not null comment('推荐时间') DATETIME"`
	Operator            string    `json:"operator" xorm:"not null comment('操作员') VARCHAR(50)"`
	RecommendCategoryid int       `json:"recommend_categoryid" xorm:"not null TINYINT(4)"`
}
