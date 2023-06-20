package models

type VideoPlayCommit struct {
	Id      int64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	VideoId int64 `json:"video_id" xorm:"not null default 0 comment('视频ID') index BIGINT(20)"`
	IsDel   int   `json:"is_del" xorm:"not null default 0 comment('0未删除，1已删除【软删除被JAVA读取过的数据】') TINYINT(1)"`
}
