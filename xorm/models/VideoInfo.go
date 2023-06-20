package models

import (
	"time"
)

type Videoinfo struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Rid               int       `json:"rid" xorm:"comment('课件id') index INT(11)"`
	Userid            string    `json:"userId" xorm:"index VARCHAR(350)"`
	Videoid           string    `json:"videoId" xorm:"comment('视频id不带老师') index VARCHAR(350)"`
	VideoidTeacher    string    `json:"videoId_teacher" xorm:"comment('带老师的视频id') VARCHAR(350)"`
	Logdate           time.Time `json:"logDate" xorm:"comment('第三方处理完成时间') index DATETIME"`
	Bjyvideoid        int       `json:"BjyVideoId" xorm:"comment('百家云视频信息') index INT(11)"`
	BjyvideoidTeacher int       `json:"BjyVideoId_teacher" xorm:"comment('百家云带教师的id') index INT(11)"`
	Isdeal            int       `json:"isDeal" xorm:"comment('是否已处理的数据') INT(255)"`
}
