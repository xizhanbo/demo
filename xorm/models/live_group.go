package models

import (
	"time"
)

type LiveGroup struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Bjyroomid   string    `json:"bjyRoomId" xorm:"comment('百家云房间ID') index VARCHAR(20)"`
	Groupid     int       `json:"groupId" xorm:"comment('房间分组ID') INT(10)"`
	Studentcode string    `json:"studentCode" xorm:"comment('分组学生邀请码') index VARCHAR(10)"`
	Userrole    int       `json:"userRole" xorm:"comment('用户角色') INT(11)"`
	Url         string    `json:"url" xorm:"comment('角色听课链接') VARCHAR(500)"`
	Status      int       `json:"status" xorm:"default 0 comment('分组状态 1:已使用 0:未使用') index INT(10)"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"comment('更新时间') DATETIME"`
	CreatedAt   time.Time `json:"created_at" xorm:"comment('创建时间') DATETIME"`
	Admincode   string    `json:"adminCode" xorm:"comment('助教码') VARCHAR(10)"`
	Followname  string    `json:"followName" xorm:"comment('跟进人用户名') VARCHAR(50)"`
	Realname    string    `json:"realName" xorm:"comment('真实姓名') VARCHAR(50)"`
}
