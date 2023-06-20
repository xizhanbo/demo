package models

type CoursewareExt struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	VideoId    string `json:"video_id" xorm:"comment('点播视频id') index VARCHAR(20)"`
	RoomId     string `json:"room_id" xorm:"comment('百家云房间id') index VARCHAR(20)"`
	SessionId  string `json:"session_id" xorm:"comment('百家云长期房间的序列号，普通房间无此参数') index VARCHAR(20)"`
	TotalSize  string `json:"total_size" xorm:"comment('视频大小，单位是字节') index VARCHAR(20)"`
	TimeLength string `json:"time_length" xorm:"default '' comment('视频时长（秒）') index VARCHAR(255)"`
}
