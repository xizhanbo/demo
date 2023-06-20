package models

type CourseIntroductionVideo struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	ClassId       int    `json:"class_id" xorm:"not null comment('课程id') index INT(11)"`
	Title         string `json:"title" xorm:"not null comment('课件标题') index VARCHAR(200)"`
	CourseVideoId string `json:"course_video_id" xorm:"not null comment('点播视频id') index VARCHAR(200)"`
	BjyAccount    string `json:"bjy_account" xorm:"default '' comment('百家云帐号') index VARCHAR(150)"`
	CoverPhoto    string `json:"cover_photo" xorm:"default '' comment('封面图片') VARCHAR(200)"`
	Status        int    `json:"status" xorm:"default 0 comment('0未删除1已删除') TINYINT(1)"`
	FileSize      int64  `json:"file_size" xorm:"default 0 comment('点播课件大小') index BIGINT(20)"`
}
