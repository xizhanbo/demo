package models

import (
	"time"
)

type CoursewareInfo struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	CoursewareId int    `json:"courseware_id" xorm:"not null comment('课件id') INT(11)"`
	VideoType    int    `json:"video_type" xorm:"default 0 comment('课件格式 1.mp4 2.flv 3.m3u8') TINYINT(1)"`
	Size         string `json:"size" xorm:"comment('课件大小') CHAR(100)"`
	Definition   string `json:"definition" xorm:"not null comment('清晰度 
low/std/high/super/1080p 
流畅/标清/高清/超清/1080P') CHAR(50)"`
	Length     string    `json:"length" xorm:"comment('视频时长 单位：秒') CHAR(100)"`
	PrefaceUrl string    `json:"preface_url" xorm:"comment('封面图片地址') VARCHAR(255)"`
	CreatedAt  time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"DATETIME"`
}
