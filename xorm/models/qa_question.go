package models

import (
	"time"
)

type QaQuestion struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Type         int       `json:"type" xorm:"not null comment('问题类型') index TINYINT(3)"`
	Content      string    `json:"content" xorm:"not null comment('问题内容') TEXT"`
	Lessionid    int       `json:"lessionid" xorm:"not null comment('课件ID') index INT(10)"`
	Point        int       `json:"point" xorm:"not null default 0 INT(10)"`
	Createtime   time.Time `json:"createtime" xorm:"not null comment('创建时间') index DATETIME"`
	Userid       int       `json:"userid" xorm:"not null comment('创建人ID') INT(10)"`
	Nickname     string    `json:"nickname" xorm:"not null default '' comment('创建人用户名') VARCHAR(20)"`
	LastUserid   int       `json:"last_userid" xorm:"not null default 0 comment('最后回复人ID') INT(10)"`
	LastNickname string    `json:"last_nickname" xorm:"not null default '' comment('最后回复人昵称') VARCHAR(20)"`
	ReplyNum     int       `json:"reply_num" xorm:"not null default 0 comment('回答数量') INT(10)"`
	BestReply    int       `json:"best_reply" xorm:"not null default 0 TINYINT(1)"`
	SelfReply    int       `json:"self_reply" xorm:"not null default 0 TINYINT(1)"`
	BestAdminSet int       `json:"best_admin_set" xorm:"not null default 0 TINYINT(1)"`
	Status       int       `json:"status" xorm:"not null default 1 comment('1正常,2取消,3删除') index TINYINT(1)"`
	Updatetime   time.Time `json:"updatetime" xorm:"comment('最后修改时间') DATETIME"`
	Mark         int       `json:"mark" xorm:"INT(10)"`
	Terminal     int       `json:"terminal" xorm:"default 3 comment('终端: 1:安卓: 2:苹果 3:pc 4:安卓Pad 5:ipad 7:M站 8:面库;') INT(10)"`
}
