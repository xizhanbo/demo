package models

import (
	"time"
)

type StatisticsUserLiveBroadcastLearningForGroup struct {
	Id                 int64     `json:"id" xorm:"pk autoincr comment('ID') BIGINT(11)"`
	BjyUserId          int64     `json:"bjy_user_id" xorm:"comment('百家云用户ID') BIGINT(11)"`
	UserId             int       `json:"user_id" xorm:"comment('在线用户ID') INT(11)"`
	NickName           string    `json:"nick_name" xorm:"comment('昵称') VARCHAR(50)"`
	LiveDate           time.Time `json:"live_date" xorm:"comment('日期') DATETIME"`
	UserRole           int       `json:"user_role" xorm:"comment('角色 0:学生 1:老师 2:助教') TINYINT(4)"`
	FirstTime          time.Time `json:"first_time" xorm:"comment('最早进入教室时间') DATETIME"`
	LastTime           time.Time `json:"last_time" xorm:"comment('最晚离开教室时间') DATETIME"`
	FirstHeartbeatTime time.Time `json:"first_heartbeat_time" xorm:"comment('上课状态下，最早在教室的时间') DATETIME"`
	LastHeartbeatTime  time.Time `json:"last_heartbeat_time" xorm:"comment('上课状态下，最晚在教室的时间') DATETIME"`
	ActualListenTime   int       `json:"actual_listen_time" xorm:"comment('实际听课时长,单位（秒）') INT(11)"`
	UserIp             string    `json:"user_ip" xorm:"comment('用户IP') VARCHAR(50)"`
	Area               string    `json:"area" xorm:"comment('用户所属省份') VARCHAR(50)"`
	City               string    `json:"city" xorm:"comment('用户所属城市') VARCHAR(50)"`
	NetworkOperator    string    `json:"network_operator" xorm:"comment('使用的网络运营商') VARCHAR(50)"`
	GroupId            int       `json:"group_id" xorm:"default 0 comment('在教室中的分组') INT(10)"`
	ClientType         int       `json:"client_type" xorm:"comment('0:PC网页 1:pc客户端 2:m站 3:ios 4:android 5:mac客户端 6:微信小程序') TINYINT(4)"`
	RoomId             string    `json:"room_id" xorm:"not null comment('百家云房间ID') index VARCHAR(50)"`
	CreateAt           time.Time `json:"create_at" xorm:"comment('创建时间') DATETIME"`
	UpdateAt           time.Time `json:"update_at" xorm:"comment('更新时间') DATETIME"`
}
