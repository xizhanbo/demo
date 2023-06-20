package models

import (
	"time"
)

type StatisticsUserLearningTimes struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	VideoType  int       `json:"video_type" xorm:"comment('1直播 2回放 3录播') INT(11)"`
	DataId     string    `json:"data_id" xorm:"comment('当 video_type =1 或者 video_type = 2 时，存 roomId 当 video_type = 3 时，存 videoId') VARCHAR(50)"`
	SessionId  string    `json:"session_id" xorm:"comment('百家云 session_id') VARCHAR(50)"`
	UserId     int64     `json:"user_id" xorm:"comment('PHP 这边 Cl_user 表 UserID') BIGINT(20)"`
	Source     int       `json:"source" xorm:"comment('终端类型 4 手机网页 5 PC 网页 6 APP 内嵌网页 7 小程序 (回放时才有)') INT(11)"`
	ActionType int       `json:"action_type" xorm:"comment('动作类型，1开始/进入房间 2结束/退出房间 （每次直播进出会更新此字段）') INT(11)"`
	ActiveTime time.Time `json:"active_time" xorm:"comment('动作触发事件，格式：YYYY-mm-dd HH:ii:ss (回放为更新时间)') DATETIME"`
	CreateTime time.Time `json:"create_time" xorm:"comment('数据插入事件，格式：YYYY-mm-dd HH:ii:ss') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	WatchTime  int       `json:"watch_time" xorm:"default 0 comment('视频最长看到的位置 单位 秒
如果 video_type = 1 是直播间总时长
如果 video_type = 2 或者 video_type = 3 是回放观看到的最长时间') INT(10)"`
	Duration int `json:"duration" xorm:"default 0 comment('video_type = 2 或者 video_type = 3 视频总时长') INT(10)"`
}
