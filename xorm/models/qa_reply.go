package models

import (
	"time"
)

type QaReply struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Qid          int       `json:"qid" xorm:"not null comment('问题ID') index INT(10)"`
	ReplyContent string    `json:"reply_content" xorm:"not null comment('回答内容') TEXT"`
	Userid       int       `json:"userid" xorm:"not null comment('创建用户ID') INT(10)"`
	Nickname     string    `json:"nickname" xorm:"not null default '' comment('昵称') VARCHAR(20)"`
	Createtime   time.Time `json:"createtime" xorm:"not null comment('创建时间') DATETIME"`
	Point        int       `json:"point" xorm:"not null default 0 comment('点赞数') INT(10)"`
	Vote         int       `json:"vote" xorm:"not null default 0 comment('点赞数') INT(10)"`
	BestReply    int       `json:"best_reply" xorm:"not null default 0 comment('最佳回复') TINYINT(1)"`
	SelfReply    int       `json:"self_reply" xorm:"not null default 0 TINYINT(1)"`
	Status       int       `json:"status" xorm:"not null default 1 index TINYINT(1)"`
	Mark         int       `json:"mark" xorm:"default 0 INT(10)"`
	Terminal     int       `json:"terminal" xorm:"default 3 comment('终端: 1:安卓: 2:苹果 3:pc 4:安卓Pad 5:ipad 7:M站 8:面库;') INT(10)"`
	ParentId     int       `json:"parent_id" xorm:"default 0 comment('回复ID') INT(10)"`
	Type         int       `json:"type" xorm:"not null default 0 comment('0正常回复，1语音回复') TINYINT(1)"`
	BjyvideoId   int       `json:"bjyvideo_id" xorm:"comment('百家云音频ID') INT(10)"`
	IsRole       int       `json:"is_role" xorm:"not null default 0 comment('0学员，1教师') TINYINT(1)"`
	Duration     int       `json:"duration" xorm:"default 0 comment('音频时长') INT(10)"`
	IsShow       int       `json:"is_show" xorm:"not null default 1 comment('0:转码不成功不展示, 1:转码成功展示') TINYINT(1)"`
}
