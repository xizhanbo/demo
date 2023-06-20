package models

import (
	"time"
)

type GkjtVideoinfoTest struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Aid        int       `json:"aid" xorm:"index INT(11)"`
	Userid     string    `json:"userId" xorm:"index VARCHAR(350)"`
	Videoid    string    `json:"videoId" xorm:"index VARCHAR(350)"`
	Logdate    time.Time `json:"logDate" xorm:"comment('记录下载完成时间') DATETIME"`
	Bjyvideoid int       `json:"BjyVideoId" xorm:"comment('百家云视频信息') INT(11)"`
}
